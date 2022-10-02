https://www.cnblogs.com/rongfengliang/p/13922297.html

https://github.com/willnorris/imageproxy
### 1、准备imageproxy和minio容器
### 2、通过imageproxy访问minio图片资源
```
http://localhost:8080/300/http://10.0.115.108:9000/photos/Snip20220907_8.png
http://localhost:8080: imageproxy服务器
300: 300px正方形
http://10.0.115.108:9000/photos/Snip20220907_8.png: 图片地址
```
### 3、修改host
```
10.0.115.108       lark-minio.com
10.0.115.108      bucket.lark-minio.com
```