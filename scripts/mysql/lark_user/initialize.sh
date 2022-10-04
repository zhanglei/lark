#!/bin/bash

MYSQL_USER=${MYSQL_USER:-root}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-lark2022}
MYSQL_HOST="127.0.0.1"
MYSQL_PORT=3306
MYSQL_DB="lark_user"
SCRIPT_PATH=$(cd $(dirname $0);pwd)
SQL_FILE="/sqls"

folder=$SCRIPT_PATH$SQL_FILE
for file in ${folder}/*
do
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} < ${file}
done

# 测试数据
for i in {1..10000};
do
  INSERT="INSERT INTO users ( uid, lark_id ) VALUES ( ${i}, ${i} );"
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} -e "$INSERT"
done


# 测试数据
for i in {1..10000};
do
  INSERT="INSERT INTO chat_members
          ( chat_id, uid, display_name, avatar_url, mute, platform,server_id, settings )
          VALUES
          ( 3333336666669999990, ${i},CONCAT('name:',${i}),CONCAT('avatar_url',${i}),0,1,1, CONCAT('{"uid":',${i},',','"receive":', true,',','"platforms":', ' [1,2]',\"}\"));"
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -D${MYSQL_DB} -e "$INSERT"
done