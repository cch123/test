#!/bin/sh

servers=("" "" "" "")
yesterday=`date +%Y%m%d --date="-1 day"`
today=`date '+%Y%m%d'`
day=$1
if [ ! $day ];
then
    day=$yesterday
fi
path=""


for var in ${servers[*]}
do

scp -l 30000  $var:$path$day ./$var.$day


expect << EOF
set timeout -1
spawn scp  -P 22 $var.$day root@ip:/path/

expect "password"
send ""
expect eof
EOF

rm -f $var.$day

done
echo $day >> history
