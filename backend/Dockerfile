# Backend Dockerfile
FROM golang:1.26-alpine

WORKDIR /app

# 设置 GOPROXY
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
