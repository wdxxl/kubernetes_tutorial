https://juejin.im/post/6844903998487527432

```
rm -rf /Users/wdxxl/tmp/gitlab
mkdir /Users/wdxxl/tmp/gitlab

docker-compose up
```

其他
```
docker run -d \
    --hostname gitlab \
    --name gitlab \
    --restart always \
    --publish 10022:22 --publish 10080:80 --publish 10443:443 \
    --volume /Users/wdxxl/tmp/gitlab/gitlab_data:/var/opt/gitlab \
    --volume /Users/wdxxl/tmp/gitlab/gitlab_logs:/var/log/gitlab \
    --volume /Users/wdxxl/tmp/gitlab/gitlab_config:/etc/gitlab \
    gitlab/gitlab-ce:latest
```