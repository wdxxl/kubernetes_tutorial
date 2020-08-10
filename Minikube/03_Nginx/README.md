### Nignx

```
docker build -t wdxxl/nginx .
docker images
docker run -p 80:80 -d wdxxl/nginx

docker stop wdxxl/nginx
docker rm wdxxl/nginx
```