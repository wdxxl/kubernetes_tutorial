

代码设置 & /metrics && /debug/pprof
```
cd go_sample
go mod init go_sample
go mod tidy
go mod vendor
docker-compose -f go_sample/docker-compose.yml up -d
```

Prometheus & Grafana & AlertManager 设置
```
docker-compose -f prometheus-server/docker-compose.yml up -d
```

```
docker-compose -f go_sample/docker-compose.yml -f prometheus-server/docker-compose.yml up -d
```

TODO 好多其他的配置和设置 验证Golang代码如何更好的配置 监控相关的matrics等
