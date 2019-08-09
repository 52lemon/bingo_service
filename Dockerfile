FROM golang:latest

# 设置工作空间
WORKDIR /go/src/bingo_service
#　拷贝项目
COPY . .
# 设置启用 mod
ENV GO111MODULE on
# 设置代理
ENV GOPROXY https://goproxy.io
# 下载依赖
RUN go mod download
# 编译
RUN go build .
# 暴露端口
EXPOSE 9000
# 启动服务
ENTRYPOINT ["./bingo_service"]
