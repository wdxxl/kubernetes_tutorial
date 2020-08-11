
OpenResty hello-world 例子

```
docker pull openresty/openresty:1.9.15.1-trusty

cd /Users/wangkexue/wdxxl/kubernetes_tutorial/Docker/01_OpenResty
docker run -d --name=openresty -p 8080:80 -v $PWD/config/nginx.conf:/usr/local/openresty/nginx/conf/nginx.conf:ro -v $PWD/logs:/usr/local/openresty/nginx/logs openresty/openresty:1.9.15.1-trusty

curl 127.0.0.1:8080

docker stop openresty && docker rm openresty
```