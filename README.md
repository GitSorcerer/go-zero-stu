# go zero stu

## goctl
https://github.com/zeromicro/go-zero

### 创建demo

```shell
go get -u github.com/zeromicro/go-zero

goctl api new greet
cd greet
go mod init
go mod tidy
go run greet.go -f etc/greet-api.yaml
```

### goctl生成api

```shell
goctl api go -api *.api -dir ./ --style=goZero

goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero

goctl env check -i -f
```

### goctl指定生成api

```shell
goctl api -o blog.api
```

### rpc 服务

生成 user.proto 文件

```shell
goctl rpc template -o user.proto
```

## docker 部署

```shell
# 到模块目录
cd user-grpc 

# 生成Dockerfile
goctl  docker -go user.go 

# 回到根目录 生成镜像
docker build -t user-grpc:v1 -f user-grpc/Dockerfile  . 

docker run --rm -it -p 8080:8080 user-grpc:v1
```

## 其他

GRPC 四种模式：流拦截器。。。

go get -u github.com/pkg/errors
