# deviceservice
# golang编译环境的镜像，用于golang环境编译
FROM golang:alpine AS builder

## ADD 添加文件夹的到容器中
ADD ./ /go/src/deviceservice/
#制作镜像的时候运行 设置 go代理， 设置go mod 模式
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on

## 切换 work 目录
WORKDIR /go/src/deviceservice/

## 配置 cgo 编译环境
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine AS runner

COPY --from=0 /go/src/deviceservice/cmd/deviceservice .
#EXPOSE 9107

CMD ./deviceservice -c /etc/golang/deviceservice.toml --switch=false
