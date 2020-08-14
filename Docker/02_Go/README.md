
Golang(Echo) x docker-composeでホットリロード用いた開発
https://qiita.com/y-ohgi/items/671de11f094f72a058b1

怎样构建 Golang Dockerfiles？
https://www.infoq.cn/article/cZRp8lJTtdgxaNfYcRTw

Step 1: 本地测试
```
go mod init 02_Go
写代码 - main.go
export GOPROXY=https://goproxy.cn,direct
go mod tidy --获取以来包

手工运行测试
curl localhost:1323
```

Step 2: Docker
```
go mod vendor  -- image下载太难了环境的问题，所以需要用mod转vendor
Dockerfile 编写
docker build -t myapp .

docker run -p 1323:1323 -d --name myapp myapp
curl localhost:1323
docker stop myapp && docker rm myapp
```

Step 3: Docker Compose
```
docker-compose up
curl localhost:1323
```

Step 4: 清理一些没用的image
```
docker images | grep none | awk '{ print $3; }' | xargs docker rmi
```

Step 5: Docker关联的Image大小
```
wdxxl@wangkexues-MacBook-Pro 02_Go % docker images
REPOSITORY                                        TAG                         IMAGE ID            CREATED             SIZE
myapp                                             latest                      ee2b54b7b5db        3 seconds ago       20.8MB
<none>                                            <none>                      00ea73ca8e2f        4 seconds ago       425MB
golang                                            1.15-alpine                 1a87ceb1ace5        2 days ago          402MB
alpine                                            latest                      a24bb4013296        2 months ago        5.57MB
```