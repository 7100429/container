FROM golang:alpine

WORKDIR /home/container/student

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY / /home/container/student

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build -o student

EXPOSE 9000

ENTRYPOINT ["./student"]
