# aws-logsearch

Search AWS CloudWatch logs all at once on the command line. This uses the aws [sdk-for-go](https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#pkg-overview). See [Configuring Credentials](https://github.com/aws/aws-sdk-go/#configuring-credentials) to configure your AWS secrets. 

# Installation

```
go get -u github.com/endgameinc/aws-logsearch
```

# Usage

```
# ~/go/bin/aws-logsearch

Usage: ~/go/bin/aws-logsearch [options]

Options:
	-p, --pattern      search pattern
	-g, --group        search specific log group
	-l, --list         list log groups
	-s, --starttime    start time, e.g. 2018-11-05 14:45:03
	-e, --endtime      end time, e.g. 2018-11-06 00:03:43
	-c, --count        print first count log matches only
	-r, --region       aws region (default us-east-1)
```

To search a log group that matches the substring "group1" for "login" and only print the first 5 results, run the following command:

```
# ~/go/bin/aws-logsearch -p "login" -g "group001" -c 5
Found Log Group: /group001/logs
[/group001/logs] [2018-10-08 13:02:47.357 +0000 UTC] [endgame.log] Oct  8 13:02:47 ip-172-16-100-174 app[47690]: INFO 200 POST /api/auth/login (127.0.0.1) 45.02ms
[/group001/logs] [2018-10-08 14.35.07.357 +0000 UTC] [endgame.log] Oct  8 14:35:07 ip-172-16-100-174 app[47683]: INFO 200 POST /api/auth/login (127.0.0.1) 43.23ms
[/group001/logs] [2018-10-08 21:28:00.357 +0000 UTC] [endgame.log] Oct  8 21:28:00 ip-172-16-100-174 app[47690]: INFO 200 POST /api/auth/login (127.0.0.1) 43.90ms
[/group001/logs] [2018-10-08 21:38:11.357 +0000 UTC] [endgame.log] Oct  8 21:38:11 ip-172-16-100-174 app[47683]: INFO 200 POST /api/auth/login (127.0.0.1) 43.46ms
[/group001/logs] [2018-10-08 21:41:05.357 +0000 UTC] [endgame.log] Oct  8 21:41:05 ip-172-16-100-174 app[47683]: WARNING 401 POST /api/auth/login (127.0.0.1) 36.47ms
```

To search all log groups for "login" on or after November 5th, 2018 after midnight (UTC) and print only the first result, run the following command:

```
# ~/go/bin/aws-logsearch -p "login" -s "2018-11-05 00:00:00" -c 1
[/group001/logs] [2018-11-08 16:11:09.357 +0000 UTC] [nginx-access.log] 127.0.0.1 - - [08/Nov/2018:16:10:08 +0000] "POST /api/v1/auth/login HTTP/1.1" 200 1250 "-" "python-requests/2.19.1" "-" - 0.314 38
[/group002/logs] [2018-11-05 00:20:36.862 +0000 UTC] [endgamelog] Nov  5 00:20:36 ip-172-16-100-252 app[27954]: INFO 200 POST /api/v1/auth/login/ (127.0.0.1) 44.07ms
[/group003/logs] [2018-11-05 00:13:57.877 +0000 UTC] [endgamelog] Nov  5 00:13:57 ip-172-16-100-178 app[10717]: INFO 200 POST /api/v1/auth/login/ (127.0.0.1) 44.71ms
[/group004/logs] [2018-11-05 00:00:25.85 +0000 UTC] [endgame.log] Nov  5 00:00:25 ip-172-16-100-97 app[20750]: INFO 200 POST /api/v1/auth/login/ (127.0.0.1) 44.67ms
[/group005/logs] [2018-11-05 00:13:57.838 +0000 UTC] [endgame.log] Nov  5 00:13:57 ip-172-16-100-97 app[12580]: INFO 200 POST /api/v1/auth/login/ (127.0.0.1) 44.79ms
[/group006/logs] [2018-11-05 00:01:02.318 +0000 UTC] [endgame.log] Nov  5 00:01:02 ip-172-16-100-176 app[24760]: INFO 200 POST /api/v1/auth/login/ (127.0.0.1) 45.19ms
[/group007/logs] [2018-11-05 00:00:40.661 +0000 UTC] [endgame.log] Nov  5 00:00:40 ip-172-31-86-221.ec2.internal app[114811]: INFO 200 POST /api/auth/login (127.0.0.1) 104.82ms
[/group009/logs] [2018-11-26 22:30:32.405 +0000 UTC] [nginx-access.log] 127.0.0.1 - - [26/Nov/2018:22:29:18 +0000] "POST /api/v1/auth/login HTTP/1.1" 200 1250 "-" "python-requests/2.19.1" "-" - 0.311 29
[/group010/logs] [2018-11-21 17:59:50.867 +0000 UTC] [nginx-access.log] 127.0.0.1 - - [21/Nov/2018:17:58:35 +0000] "POST /api/v1/auth/login HTTP/1.1" 200 1250 "-" "python-requests/2.19.1" "-" - 0.329 38
[/group011/logs] [2018-11-27 19:36:00.494 +0000 UTC] [nginx-access.log] 127.0.0.1 - - [27/Nov/2018:19:34:59 +0000] "POST /api/v1/auth/login HTTP/1.1" 200 1250 "-" "python-requests/2.19.1" "-" - 0.315 37
[/group012/logs] [2018-11-05 00:01:12.046 +0000 UTC] [endgame.log] Nov  5 00:01:11 ip-172-16-100-206 app[2088]: INFO 200 POST /api/v1/auth/login/ (127.0.0.1) 44.45ms
[/group013/logs] [2018-11-26 16:22:35.52 +0000 UTC] [nginx-access.log] 127.0.0.1 - - [26/Nov/2018:16:21:14 +0000] "POST /api/v1/auth/login HTTP/1.1" 200 1250 "-" "python-requests/2.19.1" "-" - 0.365 66
[/group014/logs] [2018-11-05 00:00:07.077 +0000 UTC] [endgame.log] Nov  5 00:00:06 ip-172-16-100-174 app[69249]: INFO 200 POST /api/auth/login (127.0.0.1) 43.06ms
[/group015/logs] [2018-11-21 18:50:43.437 +0000 UTC] [nginx-access.log] 127.0.0.1 - - [21/Nov/2018:18:49:43 +0000] "POST /api/v1/auth/login HTTP/1.1" 200 1250 "-" "python-requests/2.19.1" "-" - 0.329 41
[/group016/logs] [2018-11-16 18:15:49.196 +0000 UTC] [nginx-access.log] 127.0.0.1 - - [16/Nov/2018:18:14:46 +0000] "POST /api/v1/auth/login HTTP/1.1" 200 1250 "-" "python-requests/2.19.1" "-" - 0.327 43
[/group017/logs] [2018-11-16 18:37:06.362 +0000 UTC] [endgame.log] Nov 16 18:36:05 ip-172-16-100-226 app[9145]: INFO 200 POST /api/v1/auth/login (127.0.0.1) 313.09ms
[/group019/logs] [2018-11-05 15:43:20.041 +0000 UTC] [endgame.log] Nov  5 15:43:19 ip-172-16-100-104 app[32409]: INFO 200 POST /api/auth/login (127.0.0.1) 46.22ms
[/group020/logs] [2018-11-07 23:20:36.402 +0000 UTC] [endgame.log] Nov  7 23:19:34 ip-172-16-100-120 app[8214]: INFO 200 POST /api/v1/auth/login (127.0.0.1) 325.40ms
[/group021/logs] [2018-11-05 12:26:30.92 +0000 UTC] [endgame.log] Nov  5 12:26:30 ip-172-16-100-36 app[50664]: INFO 200 POST /api/auth/login (127.0.0.1) 44.26ms
[/group022/logs] [2018-11-06 15:19:36.358 +0000 UTC] [endgame.log] Nov  6 15:19:36 ip-172-16-100-6 app[4007]: WARNING 401 POST /api/auth/login (127.0.0.1) 334.18ms
[/group023/logs] [2018-11-05 14:33:31.492 +0000 UTC] [endgame.log] Nov  5 14:33:31 ip-172-16-100-72 app[25229]: INFO 200 POST /api/auth/login (127.0.0.1) 44.24ms
[/group024/logs] [2018-11-05 14:27:00.391 +0000 UTC] [endgame.log] Nov  5 14:27:00 ip-172-16-100-133 app[8211]: INFO 200 POST /api/auth/login (127.0.0.1) 44.88ms
[/group025/logs] [2018-11-05 05:00:31.032 +0000 UTC] [endgame.log] Nov  5 05:00:31 ip-172-16-100-128 app[13572]: INFO 200 POST /api/v1/auth/login/ (127.0.0.1) 44.65ms
[/group026/logs] [2018-11-05 12:29:15.487 +0000 UTC] [endgame.log] Nov  5 12:29:15 ip-172-16-100-147 app[66209]: INFO 200 POST /api/v1/auth/login (127.0.0.1) 43.62ms
[/group027/logs] [2018-11-05 12:29:13.295 +0000 UTC] [endgame.log] Nov  5 12:29:13 ip-172-16-100-142 app[23130]: INFO 200 POST /api/v1/auth/login (127.0.0.1) 43.83ms
[/group028/logs] [2018-11-05 12:29:14.6 +0000 UTC] [endgame.log] Nov  5 12:29:14 ip-172-16-100-228 app[4950]: INFO 200 POST /api/v1/auth/login (127.0.0.1) 44.10ms
[/group029/logs] [2018-11-05 22:45:48.284 +0000 UTC] [endgame.log] Nov  5 22:44:42 ip-172-16-100-101 app[8325]: INFO 200 POST /api/v1/auth/login (127.0.0.1) 312.97ms
[/group032/logs] [2018-11-05 13:07:40.835 +0000 UTC] [endgame.log] Nov  5 13:07:40 ip-172-16-100-186 app[77543]: INFO 200 POST /api/auth/login (127.0.0.1) 44.67ms
[/group033/logs] [2018-11-16 00:45:39.145 +0000 UTC] [endgame.log] Nov 16 00:44:25 ip-172-16-100-32 app[42013]: INFO 200 POST /api/v1/auth/login (127.0.0.1) 347.84ms
[/group034/logs] [2018-11-05 04:39:27.531 +0000 UTC] [endgame.log] Nov  5 04:39:27 ip-172-16-100-131 app[12860]: INFO 200 POST /api/auth/login (127.0.0.1) 44.26ms
[/group035/logs] [2018-11-05 12:29:16.771 +0000 UTC] [endgame.log] Nov  5 12:29:16 ip-172-16-100-31 app[8206]: INFO 200 POST /api/v1/auth/login (127.0.0.1) 43.96ms
[/group036/logs] [2018-11-05 05:00:24.778 +0000 UTC] [endgame.log] Nov  5 05:00:24 ip-172-16-100-79 app[28815]: INFO 200 POST /api/v1/auth/login/ (127.0.0.1) 43.23ms
[/group037/logs] [2018-11-05 15:20:51.171 +0000 UTC] [endgame.log] Nov  5 15:20:51 ip-172-16-100-201 app[22039]: INFO 200 POST /api/auth/login (127.0.0.1) 44.45ms
```
