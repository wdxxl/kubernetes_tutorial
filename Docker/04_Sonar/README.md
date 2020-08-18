https://qiita.com/jca02266/items/12b7d2b515a393579d0f

https://github.com/SonarSource/docker-sonarqube/blob/master/example-compose-files/sq-with-postgres/docker-compose.yml

```
docker-compose up -d

http://localhost:9000/ 登陆 admin/admin
```

```
export SONAR_SCANNER_PATH=/Users/wdxxl/Mac_Software/sonar-scanner-4.4.0.2170-macosx
export PATH="$SONAR_SCANNER_PATH/bin:$PATH"
```

```
cd /Users/wdxxl/github.com/wdxxl/kubernetes_tutorial/Docker/04_Sonar/sample
sonar-scanner
```

