# Bookstore


gRPC&gRPC-Gateway 小练习

## bookstore 介绍

书店里有很多书架，每个书架都有自己的主题和大小，分别表示摆放的图书的主题和数量。

## 要点
1.数据库
2.proto
3.写业务逻辑
   -3.1 数据库操作
   -3.2 grpc 逻辑



## proto 文件

pb/bookstore.proto

## 生成代码
```shell
protoc -I=pb \
  --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
  bookstore.proto
```

## 添加分页功能
1.把分页的代码添加进来
2.修改 `.proto` 文件
3.重新生成代码