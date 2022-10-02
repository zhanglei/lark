#!/usr/bin/env bash
MYSQL_USERNAME="root"
MYSQL_PASSWORD="lark2022"
MYSQL_HOST="127.0.0.1"
MYSQL_PORT=13308
MYSQL_DB="lark_chat_member"

folder="mysql/chat_member"

for file in ${folder}/*
do
  mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} < ${file}
done

# 测试数据
for i in {1..10};
do
  INSERT="INSERT INTO chat_members
          ( chat_id, uid, display_name, avatar_url, mute, platform,server_id, settings )
          VALUES
          ( 3333336666669999990, ${i},CONCAT('name:',${i}),CONCAT('avatar_url',${i}),0,1,1, CONCAT('{"uid":',${i},',','"receive":', true,',','"platforms":', ' [1,2]',\"}\"));"
  mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USERNAME} -p${MYSQL_PASSWORD} -D${MYSQL_DB} -e "$INSERT"
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
