# go-zero RPC 服务


## 编写RPC

1. 编写pb服务，并生成代码

```sh
goctl rpc protoc user.proto 
--go_out=./types --go-grpc_out=./types --zrpc_out=.
```
2. 完善配置结构体和配置文件 (结构体和yaml文件，一定要对应上)
3. 完善ServiceContext
4. 完善rpc的业务逻辑


### rpc 服务测试工具

一个测试grpc服务的ui工具
https://github.com/fullstorydev/grpcui

安装：

确保你电脑上的 $GOPATH/bin 目录，被添加到环境变量里面

```bash
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```

使用 `localhost:8080` 是你rpc服务的地址
```bash
grpcui -plaintext localhost:8080
```

如果出现下面的情况,需要修改配置文件，项目模式为 dev 或者 test
```bash
 grpcui -plaintext localhost:12345
Failed to dial target host "localhost:12345": dial tcp [::1]:12345: connectex: No connection could be made because the target machine actively refused it.
```

项目模式为 dev 或者 test
```yaml
Name: user.rpc
ListenOn: 0.0.0.0:8080
Mode: dev
Etcd:
  Hosts:
  - 172.16.56.137:2379
  Key: user.rpc
Mysql:
  DataSource: root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai
CacheRedis:
  - Host: 127.0.0.1:6379
    Pass: "123456"
```