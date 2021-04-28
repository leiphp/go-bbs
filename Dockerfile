FROM alpine:latest

WORKDIR /go/src/projectname

ADD bbs /go/src/projectname/

EXPOSE 8001

CMD ["./bbs"]