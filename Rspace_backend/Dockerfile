FROM golang:1.21 as build-stage

WORKDIR /app

# 配置代理
RUN go env -w GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 9090
CMD ["./main"]
