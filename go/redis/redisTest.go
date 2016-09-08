package main

import (
	"fmt"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

var sql = `
select * from g_service_worksheet
where  create_time between '2016-09-02T00:00:00+0800' and '2016-09-03T00:00:00+0800'
and (id = 211984930 or order_id = 211984930 )
and (content like  '%紧急联系人%' or enterprice like '%紧急联系人%' )
and dict_id in (12326)
order by create_time  desc limit 0,10
where  create_time between '2016-09-02T00:00:00+0800' and '2016-09-03T00:00:00+0800'
where  create_time between '2016-09-02T00:00:00+0800' and '2016-09-03T00:00:00+0800'
where  create_time between '2016-09-02T00:00:00+0800' and '2016-09-03T00:00:00+0800'
where  create_time between '2016-09-02T00:00:00+0800' and '2016-09-03T00:00:00+0800'
where  create_time between '2016-09-02T00:00:00+0800' and '2016-09-03T00:00:00+0800'
and (id = 211984930 or order_id = 211984930 )
and (content like  '%紧急联系人%' or enterprice like '%紧急联系人%' )
and dict_id in (12326)
order by create_time  desc limit 0,10
and (id = 211984930 or order_id = 211984930 )
and (content like  '%紧急联系人%' or enterprice like '%紧急联系人%' )
and dict_id in (12326)
order by create_time  desc limit 0,10
and (id = 211984930 or order_id = 211984930 )
and (content like  '%紧急联系人%' or enterprice like '%紧急联系人%' )
and dict_id in (12326)
order by create_time  desc limit 0,10
and (id = 211984930 or order_id = 211984930 )
and (content like  '%紧急联系人%' or enterprice like '%紧急联系人%' )
and dict_id in (12326)
order by create_time  desc limit 0,10
and (id = 211984930 or order_id = 211984930 )
and (content like  '%紧急联系人%' or enterprice like '%紧急联系人%' )
and dict_id in (12326)
order by create_time  desc limit 0,10
`

func main() {
	c, err := redis.Dial("tcp", ":6379")
	defer c.Close()
	if err != nil {
		println(err)
		os.Exit(1)
	}

	fmt.Println(time.Now())
	for i := 0; i < 10000; i++ {
		//c.Send("set", fmt.Sprintf("foo%d", i), sql)
		_, _ = c.Do("get", fmt.Sprintf("foo%d", i))
	}
	fmt.Println(time.Now())
}
