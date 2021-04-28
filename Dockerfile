FROM alpine:latest

WORKDIR /go/src/projectname

#将容器外项目文件拷贝至容器中
COPY . .

#ADD bbs /go/src/projectname/

EXPOSE 8001

#CMD ["./bbs"]

#最终运行docker的命令
ENTRYPOINT  ["./bbs"]