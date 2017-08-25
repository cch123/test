package main

import (
	"fmt"

	"github.com/cch123/codisxredis"
)

func init() {
	var conf = codisxredis.Config{
		CodisAddrs: []string{"127.0.0.1:6379", "127.0.0.1:7777"},
		CodisName:  "r",
	}

	codisxredis.InitClientMap([]codisxredis.Config{conf})
}

func main() {
	c, _ := codisxredis.GetClient("r")
	c.Set("abc", "fuc", 0).Result()
	c.Set("xxx", "", 0).Result()
	c.Get("abc").Result()
	res, err := c.LRange("abc", 0, -1).Result()
	fmt.Println(res, err)

	//time.Sleep(time.Hour)

	dres, err := c.Get("ddd").Result()
	fmt.Println(dres, err)

	dres, err = c.Get("xxx").Result()
	fmt.Println(dres, err)
}
