## golang 网络相关---建立 TCP 等网络连接

### 基础

- 客户端
  - 使用 conn, err := net.Dail() 建立连接
  - 使用 conn.writer([]byte) 发送数据
- 服务端
  - 使用 listener, err := net.Listen() 建立监听
  - 使用 conn, err := listener.Accept() 等待连接的建立
  - 使用 err.Read([]byte) 读取客户端传过来的数据
  - 使用 conn.Close() 关闭连接