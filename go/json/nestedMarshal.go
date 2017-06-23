// 这个文件是给改造和重构 elasticsql 的 agg 部分做的一些准备
// functions used in this file is to refactor the agg
// code of elasticsql
package main

import "fmt"
import "encoding/json"

type H map[string]interface{}

var aggTemplate = `
{
	"aggregations" : {
		"{{fieldNameOrAlias}}": {
			"terms" : {
				"field" : "{{fieldName}}",
				"size" : {{fieldSize}}
			},
			"aggregations" : {{innerAggObj}}
		}
	}
}
`

var aggFieldAgg = `
{
	"{{fieldNameOrAlias}}" : {
		"terms" : {
			"field" : "{{fieldName}}",
			"size" : {{fieldSize}}
		}
	}
}
`

func buildFieldAggMap(fieldName string, fieldAlias string, aggSize int) H {
	return H{
		fieldAlias: H{
			"terms": H{
				"field": fieldName,
				"size":  aggSize,
			},
		},
	}
}

// 可以并列在最内部的aggregation列表中
var normalFuncAgg = `
{
	"{{funcName}}" : {
		"field" : "{{fieldName}}"
	}
}
`

//和上面的一样，不过count要特殊处理
var countFuncAgg = `
{
	"value_count" : {
		"field" : "_index"
	}
}
`

// var aggMap
// aggMap['aggs'] = aggObj
// aggObj['aggs'] = nestedAggObj

func main() {
	m := buildFieldAggMap("id", "user_id", 200)
	fmt.Println(m)
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
	m["aggregation"] = buildFieldAggMap("name", "name", 0)
	res, _ = json.Marshal(m)
	fmt.Println(string(res))
}
