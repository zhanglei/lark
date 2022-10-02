### etcd
```
etcd:
  endpoints: ["lark-etcd:12379"]
```

### auth
```
name: lark_auth_server
grpc_server:
  name: lark_auth_server
  port: 30000
  
mysql:
  address: "lark-mysql-user:13306"
  username: root
  password: lark2022
  db: lark_user
  
redis:
  address: ["lark-redis:63791"]
```

### chat_member
```
name: lark_chat_member_server
grpc_server:
  name: lark_chat_member_server
  port: 36000
  
mysql:
  address: "lark-mysql-member:13308"
  username: root
  password: lark2022
  db: lark_chat_member
  
redis:
  address: ["lark-redis:63791"]
```

### chat_msg
```
name: lark_chat_msg_server
grpc_server:
  name: lark_chat_msg_server
  port: 35000
  
mysql:
  address: "lark-mysql-msg:13307"
  username: root
  password: lark2022
  db: lark_chat_msg
  
redis:
  address: ["lark-redis:63791"]
```

### message
```
name: lark_message_server
grpc_server:
  name: lark_message_server
  port: 33000
```

### msg_gateway
```
name: lark_msg_gateway_server
grpc_server:
  name: lark_online_push_server
  port: 32000

ws_server:
  name: lark_ws_server
  port: 32001
```

### msg_history
```
name: lark_msg_store_server

mysql:
  address: "lark-mysql-msg:13307"
  username: root
  password: lark2022
  db: lark_chat_msg
  
redis:
  address: ["lark-redis:63791"]
```

### offline_push
```
name: lark_offline_push_server

mysql:
  address: "lark-mysql-user:13306"
  username: root
  password: lark2022
  db: lark_user
```

### push
```
name: lark_push_server
grpc_server:
  name: lark_push_server
  port: 34000

redis:
  address: ["lark-redis:63791"]
```

### user服务
```
name: lark_user_server
grpc_server:
  name: lark_user_server
  port: 31000

mysql:
  address: "lark-mysql-user:13306"
  username: root
  password: lark2022
  db: lark_user
  
redis:
  address: ["lark-redis:63791"]
```

### 错误编码
```
http: 10000~19999
user: 20000~29999
auth: 30000~39999
msg_gateway: 60000~69999
---ws: 60000~60999
---svc_ws: 61000~61999
---gateway: 62000~69999
message: 70000~79999
msg_history: 80000~84999
msg_hot: 85000~89999
push: 90000~99999
chat_member: 100000~109999
chat_msg: 110000~119999
request: 120000~129999
link: 130000~139999
```

### minio
```
init minio buckets

documents
photos
videos
```