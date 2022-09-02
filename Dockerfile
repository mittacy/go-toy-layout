FROM golang as builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,http://goproxy.xiaoe-tools.com,direct GO111MODULE=on

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN go clean

# 构建二进制文件命令,替换为自身程序的构建命令
RUN go build -o ./server .

FROM centos:7

WORKDIR /app

# 大型项目可只copy二进制文件main，加快构建速度与缩小镜像体积
COPY --from=builder /app/ .

# 启动命令，多行参数使用,隔开
ENTRYPOINT ["./server", "start", "http", "-c", ".env.docker", "-p", "8080", "-e", "development"]