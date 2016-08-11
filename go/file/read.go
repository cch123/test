package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

func main() {
	f, err := os.Open("merged")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer f.Close()

	queryDict := make(map[string]int)
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		buf = strings.TrimSpace(buf)

		if buf == "" {
			break
		}

		jsonBody, err := simplejson.NewJson([]byte(buf))
		if err != nil {
			fmt.Println(err)
			//可以查看查询参数有什么问题
			//fmt.Println(buf)
			continue
		}
		queryArr := jsonBody.Get("query").Get("bool").Get("must").MustArray()
		//[{"term" : {"phone" : "15980270151"}}]
		//fmt.Println(queryArr)
		for _, valInterface := range queryArr {
			valMap := valInterface.(map[string]interface{})

			for _, innerInterface := range valMap {
				innerMap, _ := innerInterface.(map[string]interface{})
				//calc key occurs
				for key := range innerMap {
					queryDict[key]++
				}
			}

		}
	}
	fmt.Println("calc result")
	for key, count := range queryDict {
		fmt.Println(key, count)
	}

}
