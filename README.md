# go-bbs
go-bbs是一个基于golang gin框架开发的个人社区论坛项目，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点，让你快速了解Gin框架的使用

## 说明
Dockerfile2是使用docker-compose.yml构建的文件  
Dockerfile是Jenkins内部构建的文件  
k8s目录配置k8s运行的文件清单  

## 普遍部署
构建镜像
docker build -t go-bbs .  
运行容器
docker run  -p 8001:8001 --name http_gobbs -d go-bbs   

## docker-composer部署
采用docker-composer部署（默认Dockerfile是Dockerfile2需要把文件改过来）  
docker-compose -f docker-compose.yml up -d  

## k8s部署
指定文件构建镜像  
docker build -f  Dockerfile2 -t 192.168.101.208/test/go-bbs:latest .  
登录私有镜像仓库  
docker login 192.168.101.208  
推送镜像到私有仓库   
docker push 192.168.101.208/test/go-bbs:latest

