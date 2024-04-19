import socket
import json

requst = {
    "id": 0,
    "param": [{"x":10, "y":20}],  # 参数要对应上Args结构体
    "reply": 0,
    "method": "ServiceA.Add"
}

client = socket.create_connection(("127.0.0.1", 9091),5)
client.sendall(json.dumps(requst).encode())

rsp = client.recv(1024)
rsp = json.loads(rsp.decode())
print(rsp)