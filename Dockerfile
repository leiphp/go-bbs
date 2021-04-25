FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

ADD main /app/main

EXPOSE 8001

CMD ["./main"]