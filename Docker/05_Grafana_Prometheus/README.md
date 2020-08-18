

代码设置
```
cd go_sample
go mod init go_sample
go mod tidy
go mod vendor
docker-compose -f go_sample/docker-compose.yml up
```

Prometheus & Grafana 设置
```
docker-compose -f prometheus-server/docker-compose.yml up
```

TODO 好多其他的配置和设置 验证Golang代码如何更好的配置 监控相关的matrics等
