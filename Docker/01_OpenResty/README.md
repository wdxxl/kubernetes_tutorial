
OpenResty hello-world 例子

```
docker pull openresty/openresty:1.9.15.1-trusty

cd /Users/wangkexue/wdxxl/kubernetes_tutorial/Docker/01_OpenResty
docker run -d --name=openresty -p 8080:80 -v $PWD/config/nginx.conf:/usr/local/openresty/nginx/conf/nginx.conf:ro -v $PWD/logs:/usr/local/openresty/nginx/logs openresty/openresty:1.9.15.1-trusty

curl 127.0.0.1:8080

docker stop openresty && docker rm openresty
```

AB结果查看
```
wangkexuedeMacBook-Pro:01_OpenResty wangkexue$ ab -c10 -n50000 http://localhost:8080/
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 5000 requests
Completed 10000 requests
Completed 15000 requests
Completed 20000 requests
Completed 25000 requests
Completed 30000 requests
Completed 35000 requests
Completed 40000 requests
Completed 45000 requests
Completed 50000 requests
Finished 50000 requests


Server Software:        openresty/1.9.15.1
Server Hostname:        localhost
Server Port:            8080

Document Path:          /
Document Length:        20 bytes

Concurrency Level:      10
Time taken for tests:   149.919 seconds
Complete requests:      50000
Failed requests:        0
Total transferred:      8400000 bytes
HTML transferred:       1000000 bytes
Requests per second:    333.51 [#/sec] (mean)
Time per request:       29.984 [ms] (mean)
Time per request:       2.998 [ms] (mean, across all concurrent requests)
Transfer rate:          54.72 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.6      0     135
Processing:     4   30  18.7     26     242
Waiting:        4   29  18.3     25     242
Total:          4   30  18.8     26     242

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     30
  75%     33
  80%     36
  90%     43
  95%     53
  98%    101
  99%    128
 100%    242 (longest request)
wangkexuedeMacBook-Pro:01_OpenResty wangkexue$ 
```