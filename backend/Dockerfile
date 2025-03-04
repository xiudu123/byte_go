# 定义基础镜像
FROM golang:1.22.5

# 声明工作目录
WORKDIR /usr/src/byte_go/backend

# 设置环境变量GOPROXY
ENV GOPROXY=https://goproxy.io,direct

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY app/frontend/go.mod app/frontend/go.sum ./app/frontend/
COPY rpc_gen rpc_gen

RUN cd app/frontend/ && go mod download && go mod verify

COPY app/frontend app/frontend

RUN cd app/frontend/ && go build -v -o /opt/byte_go/backend/frontend/server

CMD ["app"]
