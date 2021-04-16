FROM golang

#修改系统为上海时区
RUN echo "Asia/Shanghai" > /etc/timezone \
 && rm /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

#开启go mod 模式
ENV GO111MODULE on
#必须配置为0，否则docker容器中编译失败，CGO_ENABLED=0的情况下，Go采用纯静态编译，避免各种动态链接库依赖的问题
ENV CGO_ENABLED 0
#切换到工作路径，建议到/go/src 路径下，曾在将项目文件拷贝至容器时，由于配置其他项目，导致一直拷贝不成功
WORKDIR /go/src/projectname
#将容器外项目文件拷贝至容器中
COPY . .
#安装依赖
RUN go mod tidy
#编译
RUN go build