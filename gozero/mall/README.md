# go-zero RPC 服务


## 编写 API
1. 编写 .api 文件，生辰代码

```bash
goctl api go -api order.api -dir . -style=goZero
```

2. 使用 goctl + sql 文件生成 model 层代码
注意: 1.主键一定要单独定义约束
      2.sql文件中index要这样写`UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE`才能生成对应根据index查找的代码
```sql
CREATE TABLE Users(
    id int,
    username varchar(255) not null,
    primary key(id),
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
)
```
```bash
goctl model mysql ddl -src .\api\orders.sql -dir .\model -c

```

其中：-src：sql文件路径  -dir:生成的文件目录  -c:是否开启缓存相关代码




## 编写RPC

1. 编写pb服务，并生成代码

```sh
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
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
 grpcui -plaintext localhost:12345 Failed to dial target host "localhost:12345": dial tcp [::1]:12345: connectex: No connection could be made because the target machine actively refused it.
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


## 订单服务的检索接口

/api/order/search: 根据订单id查询订单信息
  -RPC--> userID -> user.GetUser

课后作业：
1. 把订单服务自己完善一下  


## go-zero 中通过RPC调用其他服务

1. 配置RPC客户端（配置结构体和yaml配置文件都要加RPC客户端客户端配置，注意：etcd的key要对应上）
2. 修改 ServiceContext （告诉生成的代码我们现在有RPC客户端了）
    - go-zero中的RPC服务会自动生成一份客户端的代码
3. 编写业务逻辑（可以直接通过RPC客户端发起RPC调用了）


## 使用Consul作为注册中心

### 服务注册
1. 修改配置 (配置结构体和yaml配置文件)
    - 引入 "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
    - 注释掉原来默认的etcd，添加Consul相关配置
2. 服务启动的时候将服务注册到Consul
    - consul.RegisterService(c.ListenOn, c.Cousul)    

### 服务发现
1. 修改配置
    - Target: consul://127.0.0.1:8500/consul-user.rpc?wait=14s
2. 程序启动时引入 import _ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"


## RPC 调用传递metadata


### 几个知识点
1. 什么是metadata? 什么样的数据应该传入metadata? 它和请求参数有什么区别
2. gRPC的拦截器: 客户端的拦截器和服务端的拦截器


### go-zero项目添加client端拦截器

order服务的search接口中添加拦截器，添加一些requestID、token、userID等数据

几个关键点：
1. 什么时候存入metadata
2. 怎么存
3. 拦截器中如何通过context传值
4. context存值取值操作

### go-zero项目添加server端拦截器
1. 拦截器怎么加，什么时候加
2. 拦截器的业务逻辑怎么写
3. 服务端拦截器如何从metadata中取值


## 错误处理

```json
{
    "code": 10001,
    "message": "内部错误"

}
```

1. 定义自定义错误格式
2. 业务代码中要按需返回自定义的错误
3. 告诉go-zero框架处理一下我们的自定义错误