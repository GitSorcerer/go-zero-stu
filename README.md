```shell
goctl api new greet
cd greet
go mod init
go mod tidy
go run greet.go -f etc/greet-api.yaml
```


goctl api go -api *.api -dir ./ --style=goZero


 goctl env check -i -f


 Api 层相关命令：

执行命令 goctl api -o blog.api, 创建 blog.api 文件。

执行命令 goctl api go -api blog.api -dir . ，生成 api 相关代码。

加参数 goctl 也可以生成其他语言的 api 层的文件，比如 java、ts 等，尝试之后发现很难用，所以不展开了。

rpc 服务
编写 proto 文件

生成 user.proto 文件
使用命令 goctl rpc template -o user.proto, 生成 user.proto 文件

user.proto 文件的作用
user.proto 的作用是用来生成 rpc 服务的相关代码。 

生成 rpc 相关代码

生成 user rpc 服务相关代码
使用命令 goctl rpc proto -src user.proto -dir . 生成 user rpc 服务的代码。

小结

rpc 服务相关命令：

使用命令 goctl rpc template -o user.proto, 生成 user.proto 文件

使用命令 goctl rpc proto -src user.proto -dir . 生成 user rpc 服务的代码。