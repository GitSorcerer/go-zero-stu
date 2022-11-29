# grpc操作db
```shell
go install github.com/Mikaelemmmm/sql2pb@latest

sql2pb -go_package ./pb -host localhost -package pb -password 123456 -port 3306 -schema go-zero -service_name userOrder -user root > userOrder.proto

goctl rpc protoc  *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero
```

jinzhu   copier.Copy()

 git clone https://GitSorcerer:XXXX@github.com/GitSorcerer/go-zero-looklook.git

 Reply from ::1: time<1ms