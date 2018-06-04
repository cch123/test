package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type msi = map[string]interface{}

var jStr = `{
	"data" : {
		"key" : {
			"val" : 1
		}
	}
}`

var ruleMap = map[string]string{
	"data.key.val": "user_id",
}

func readVal(ruleStr string, srcMap msi) interface{} {
	var res interface{}
	keyArr := strings.Split(ruleStr, ".")
	loopMap := srcMap
	fmt.Println(srcMap)
	fmt.Println(keyArr)

	leafKey := keyArr[len(keyArr)-1]
	for i := 0; i < len(keyArr)-1; i++ {
		key := keyArr[i]
		_, keyExist := loopMap[key]
		if !keyExist {
			println("field not exist aaa")
			break
		}

		fmt.Printf("%#v\n", loopMap[key])
		if childMap, ok := loopMap[key].(msi); ok {
			loopMap = childMap
		} else {
			println("field not exist bbb")
			break
		}
		fmt.Println(loopMap)
	}
	res = loopMap[leafKey]
	return res
}

func main() {
	var iter interface{}
	err := json.Unmarshal([]byte(jStr), &iter)
	println(err)
	fmt.Printf("%#v\n", iter)

	res := readVal("data.key.val", msi{
		"data": msi{
			"key": msi{
				"val": 1,
			},
		},
	})
	fmt.Println("rrr", res, reflect.TypeOf(res), reflect.ValueOf(res))
}
