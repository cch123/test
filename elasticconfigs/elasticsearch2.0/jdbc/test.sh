#!/bin/sh

curl -XDELETE 'localhost:9200/worksheets'

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
bin=${DIR}/../bin
lib=${DIR}/../lib

echo '
{
    "type" : "jdbc",
    "jdbc" : {
        "url" : "jdbc:mysql://10.10.10.107:3306/ark_callcenter",
        "user" : "root",
        "password" : "123456",
        "sql" : "select *, id as _id from g_service_worksheet",
        "index" : "worksheets",
	"type" : "service_worksheet"
    }
}
' | java \
    -cp "${lib}/*" \
    -Dlog4j.configurationFile=${bin}/log4j2.xml \
    -Xms2g -Xmx2g -Djava.awt.headless=true -XX:+UseParNewGC -XX:+UseConcMarkSweepGC -XX:CMSInitiatingOccupancyFraction=75 -XX:+UseCMSInitiatingOccupancyOnly -XX:+HeapDumpOnOutOfMemoryError -XX:+DisableExplicitGC -Dfile.encoding=UTF-8 \
    org.xbib.tools.Runner \
    org.xbib.tools.JDBCImporter


echo "sleeping while river should run..."

sleep 10

curl -XPOST 'localhost:9200/worksheets/_search?pretty' -d '
{
       "query": {
          "match_all": {
           }
       }
}'

# expected

#{
#  "took" : 15,
#  "timed_out" : false,
#  "_shards" : {
#    "total" : 5,
#    "successful" : 5,
#    "failed" : 0
#  },
#  "hits" : {
#    "total" : 1,
#    "max_score" : 1.0,
#    "hits" : [ {
#      "_index" : "mytest",
#      "_type" : "jdbc",
#      "_id" : "0",
#      "_score" : 1.0,
#      "_source":{"movie":{"title":"Krieg der Welten","overview":"Eines windigen herbstlichen Nachmittags wird der","test":1212}}
#    } ]
#  }
#}
