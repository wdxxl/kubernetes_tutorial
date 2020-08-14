
Golang(Echo) x docker-composeでホットリロード用いた開発
https://qiita.com/y-ohgi/items/671de11f094f72a058b1


```
go mod init 02_Go
写代码 - main.go
export GOPROXY=https://goproxy.cn,direct
go mod tidy

curl localhost:1323
```

```
go mod vendor image下载太难了
Dockerfile 编写
docker build -t myapp .

docker run -p 1323:1323 -d --name myapp myapp
curl localhost:1323
docker stop myapp && docker rm myapp
```

```
docker-compose up
curl localhost:1323
```
