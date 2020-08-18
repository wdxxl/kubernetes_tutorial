
Dockerfile
```
docker build -t wdxxl/nginx .    
docker run -d --name=nginx -p 8080:80 wdxxl/nginx 

docker stop nginx && docker rm nginx  
```