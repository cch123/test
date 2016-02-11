#!/bin/sh

servers=("" "" "" "")
yesterday=`date +%Y%m%d --date="-1 day"`
today=`date '+%Y%m%d'`
day=$1
if [ ! $day ];
then
    day=$yesterday
fi
path="/home/logs/x.log"


for var in ${servers[*]}
do

scp -l 30000  $var:$path$day ./$var.$day

done

