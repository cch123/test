#!bin/bash
counter=0
slCounter=100
while read line
do
   mysql -h127.0.0.1 -uroot -p123456 -e "$line"
   echo $line
   counter=$(($counter+1))
   echo $counter

   if [ $counter -eq $slCounter ]
   then
       sleep 1
       counter=0
       echo "oh yes"
   fi
done < res.txt

