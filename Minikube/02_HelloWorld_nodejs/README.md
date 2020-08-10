### Hello World

#### 1.Docker 例子 - 运行测试一下效果
```
docker build -t wdxxl/hello-world .
docker images
docker run -d --name=hello-world -p 8888:8080 wdxxl/hello-world
curl 127.0.0.1:8888

docker exec -it hello-world bash
curl 127.0.0.1:8080

docker stop hello-world
docker rm hello-world
```

#### 2. Minikube 运行 - kubectl run

https://stackoverflow.com/questions/42564058/how-to-use-local-docker-images-with-minikube
```
前期准备 - 启动
minikube start --vm-driver=virtualbox
minikube docker-env
eval $(minikube -p minikube docker-env)

运行本地的image 测试效果 - only pod
kubectl run hello-world --image=wdxxl/hello-world --image-pull-policy=Never --port=9999
kubectl get pods,deployments,service
登录pod测试
kubectl exec -it hello-world sh
curl 127.0.0.1:8080

kubectl delete pods hello-world
```

#### 3. Minikube 运行 - kubectl create
https://www.edureka.co/community/17081/using-a-local-image-to-create-a-pod-in-k8s
```
docker login --username=wdxxl
docker push wdxxl/hello-world

kubectl create -f hello_world.yml
kubectl get pods,deployments,service

kubectl expose deployment hello-world --type=NodePort --port=8080

curl $(minikube service hello-world --url)

minikube dashboard
```

