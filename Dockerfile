FROM golang:latest

WORKDIR /home/container/student

COPY / /home/container/student

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build -o student

ENTRYPOINT ["./student"]
