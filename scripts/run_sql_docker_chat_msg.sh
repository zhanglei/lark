#!/usr/bin/env bash
MYSQL_USERNAME="root"
MYSQL_PASSWORD="lark2022"
MYSQL_HOST="127.0.0.1"
MYSQL_PORT=13307
MYSQL_DB="lark_chat_msg"

folder="mysql/chat_msg"

for file in ${folder}/*
do
  mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} < ${file}
done

<<xxxx

mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} < ${file}

mysql -h127.0.0.1 -P3306 -uroot -p123456 -Dsdb
#参数
-h:host主机
-P:port端口
-u:user用户名
-p:password密码
-D:database数据库

xxxx
