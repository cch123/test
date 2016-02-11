#!/bin/sh

servers=("aliyun")
yesterday=`date +%Y-%m-%d`
path="/home/logs/catalina.out."


for var in ${servers[*]}
do

scp -l 30000 -P 36000 $var:$path$yesterday ./$var.$yesterday

expect <<EOF
spawn scp -l 20000 -P 22 $var.$yesterday root@ip:/var/ftp/pub

expect "password"
send "mypassword"
expect eof
EOF

#rm -f $var.$yesterday

sleep 10

done

