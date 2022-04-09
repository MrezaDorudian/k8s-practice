# kubernetes-learning
Learning K8s for cloud-computing course

## Part 1
mostly learning about building a Dockerfile
at ./curl-alpine there is a Dockerfile that we can build a modified image of alpine which curl installed on it.

to get the image:
- make a new image:  
`cd ./curl-alpine`  
`docker build -t imageName:tag .`
- get from dockerhub:  
`docker pull elmo4679/curl-alpine`

to run it once and get a bash (second method):  
`docker run -it --rm elmo4679/curl-alpine`

now there is a bash and an installed curl  
`bash-5.1# curl google.com`

## Part 2
developing a simple http server with go and echo framework that returns weather situation of given city, running on port 8080 as default
`curl http://localhost:8080/tehran`
`{"hostname":"C:\\Users\\mrdor@ThisPC","temperature":12,"weather_descriptions":["Clear"],"wind_speed"19"humidity":14"feelslike":11`  
then dockerize it with Docker file at ./Dockerfile

to get the image:
- make a new image:
  `docker build -t imageName:tag .`
- get from dockerhub:  
  `docker pull elmo4679/weather-app`

## Part 3
creating yaml files for k8s config-map, deployment and service

to build the project on minikube:
- first install minikube
- `minikube start`
- `kubectl apply -f ./deployment/config-map.yaml`
- `kubectl apply -f ./deployment/deployment.yaml`
- `kubectl apply -f ./deployment/service.yaml`
- `kubectl get pods`

now you can see the pods running on minikube

## Part 4
now to test the service, we make another pod with curl-alpine image and run it on minikube
then with the curl command we can test it.  
get ip address or name of the service:  
`kubectl get svc -o wide`
then `kubectl run -it --rm curl-test --image=elmo4679/curl-alpine /bin/bash`

now with the bash we can test the service:  
`bash-5.1# curl http://{IP address result of previous command}:8080/tehran`  
or  
`bash-5.1# curl weather-app-service:8080/tehran`



