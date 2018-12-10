package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var usageStr = `
Usage: %s [options]

Options:
	-p, --pattern      search pattern
	-g, --group        search specific log group
	-l, --list         list log groups
	-s, --starttime    start time, e.g. 2018-11-05 14:45:03
	-e, --endtime      end time, e.g. 2018-11-06 00:03:43
	-c, --count        print first count log matches only
	-r, --region       aws region (default us-east-1)
`

func usage() {
	fmt.Printf("%s\n", fmt.Sprintf(usageStr, os.Args[0]))
	os.Exit(0)
}

func CaseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func printLogGroups(lgs *cloudwatchlogs.DescribeLogGroupsOutput) {
	fmt.Println(lgs)
}

func searchLogGroup(
	logs *cloudwatchlogs.CloudWatchLogs,
	name string,
	pattern string,
	token string,
	startTime int64,
	endTime int64,
	count int) {

	if count == 0 {
		return
	}

	filter := &cloudwatchlogs.FilterLogEventsInput{}
	filter.SetFilterPattern(pattern)
	filter.SetLogGroupName(name)

	if startTime != 0 {
		filter.SetStartTime(startTime)
	}

	if endTime != 0 {
		filter.SetEndTime(endTime)
	}

	if token != "" {
		filter.SetNextToken(token)
	}

	resp, err := logs.FilterLogEvents(filter)
	if err != nil {
		fmt.Println(err)
		return
	}

	eventCount := 0

	location, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, m := range resp.Events {
		if eventCount == count {
			return
		}

		msg := *m.Message
		timestamp := *m.Timestamp
		filename := *m.LogStreamName

		t := time.Unix(0, (timestamp * int64(time.Millisecond))).In(location)
		fmt.Printf("[%s] [%s] [%s] %s\n", name, t, filename, msg)

		eventCount += 1
	}

	if resp.NextToken != nil && *resp.NextToken != "" {
		searchLogGroup(logs, name, pattern, *resp.NextToken, startTime, endTime, (count - eventCount))
	}
}

func searchLogGroups(
	logs *cloudwatchlogs.CloudWatchLogs,
	lgs *cloudwatchlogs.DescribeLogGroupsOutput,
	pattern string,
	startTime int64,
	endTime int64,
	count int) {

	var wg sync.WaitGroup

	for _, v := range lgs.LogGroups {
		name := *v.LogGroupName

		wg.Add(1)
		go func() {
			defer wg.Done()
			searchLogGroup(logs, name, pattern, "", startTime, endTime, count)
		}()
	}

	wg.Wait()
}

func matchLogGroupName(lgs *cloudwatchlogs.DescribeLogGroupsOutput, group string) string {
	for _, v := range lgs.LogGroups {
		name := *v.LogGroupName
		if CaseInsensitiveContains(name, group) {
			return name
		}
	}
	return ""
}

func parseTimestampToMillis(timestamp string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return 0, err
	}
	nanos := t.UnixNano()
	return (nanos / int64(time.Millisecond)), nil
}

func main() {

	desc := "pattern"
	defaultValue := ""
	pattern := defaultValue
	flag.StringVar(&pattern, "p", defaultValue, desc)
	flag.StringVar(&pattern, "pattern", defaultValue, desc)

	desc = "list"
	list := false
	flag.BoolVar(&list, "l", false, desc)
	flag.BoolVar(&list, "list", false, desc)

	desc = "log group"
	group := ""
	flag.StringVar(&group, "g", defaultValue, desc)
	flag.StringVar(&group, "group", defaultValue, desc)

	desc = "start time, e.g. 2018-11-05 14:45"
	startTimeString := ""
	flag.StringVar(&startTimeString, "s", defaultValue, desc)
	flag.StringVar(&startTimeString, "startime", defaultValue, desc)

	desc = "end time, e.g. 2018-11-06 00:03"
	endTimeString := ""
	flag.StringVar(&endTimeString, "e", defaultValue, desc)
	flag.StringVar(&endTimeString, "endtime", defaultValue, desc)

	desc = "print first count log matches only"
	count := -1
	flag.IntVar(&count, "c", count, desc)
	flag.IntVar(&count, "count", count, desc)

	desc = "region (default us-east-1)"
	region := "us-east-1"
	flag.StringVar(&region, "r", region, desc)
	flag.StringVar(&region, "region", region, desc)

	flag.Usage = usage
	flag.Parse()

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	logs := cloudwatchlogs.New(sess)

	lgs, err := logs.DescribeLogGroups(&cloudwatchlogs.DescribeLogGroupsInput{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if list == true {
		printLogGroups(lgs)
		os.Exit(0)
	}

	startTime := int64(0)
	endTime := int64(0)

	if startTimeString != "" {
		startTime, err = parseTimestampToMillis(startTimeString)
		if err != nil {
			fmt.Printf("error parsing start timestamp: %s\n", err)
			return
		}
	}
	if endTimeString != "" {
		endTime, err = parseTimestampToMillis(endTimeString)
		if err != nil {
			fmt.Printf("error parsing end timestamp: %s\n", err)
			return
		}
	}

	if group != "" {
		matchedGroupName := matchLogGroupName(lgs, group)
		if matchedGroupName != "" {
			fmt.Printf("Found Log Group: %s\n", matchedGroupName)
			searchLogGroup(
				logs,
				matchedGroupName,
				pattern,
				"",
				startTime,
				endTime,
				count)
		}
		os.Exit(0)
	}

	if pattern == "" {
		usage()
	} else {
		searchLogGroups(logs, lgs, pattern, startTime, endTime, count)
	}
}
