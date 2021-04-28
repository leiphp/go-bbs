FROM alpine:latest

WORKDIR /go/src/projectname

EXPOSE 8001

CMD ["./bbs"]