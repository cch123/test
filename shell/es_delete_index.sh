#! /bin/sh
date=`date -d "-15 day $1" +%Y%m%d`
enddate=`date -d "-5 day $2" +%Y%m%d`

echo "------------------------------"
echo "date=$date"
echo "enddate=$enddate"
echo "------------------------------"


while [[ $date < $enddate  ]]
do
        curl -XDELETE localhost:9200/optimus-new-$date
        echo
        date=`date -d "+1 day $date" +%Y%m%d`
done
