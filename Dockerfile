FROM golang:1.19

MAINTAINER blackfish
# 仅指定镜像元数据内容
LABEL wonderful-pages=1.0.0

# 添加本地仓库作为镜像的源代码和依赖项 目标路径需要使用绝对路径
COPY . /wonderful-pages
WORKDIR /wonderful-pages

# 配置代理 从代理拉取go依赖
RUN export GOPROXY=https://mirrors.aliyun.com/goproxy/ && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wonderful-pages wonderful-pages
EXPOSE 9090
CMD ["./wonderful-pages","conf/config.ini"]

# docker build -t wonderful-pages .