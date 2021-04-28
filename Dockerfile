FROM alpine:latest

WORKDIR /go/src/projectname

EXPOSE 8001

#CMD ["./bbs"]
#最终运行docker的命令
#ENTRYPOINT  ["./bbs"]
CMD ["/go/src/projectname/bbs"]