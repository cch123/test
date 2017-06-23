// 这个文件是给改造和重构 elasticsql 的 agg 部分做的一些准备
// functions used in this file is to refactor the agg
// code of elasticsql
package main

import "fmt"
import "encoding/json"

type H map[string]interface{}

func buildGroupByField(fieldName string, fieldAlias string, aggSize int) H {
	return H{
		fieldAlias: H{
			"terms": H{
				"field": fieldName,
				"size":  aggSize,
			},
		},
	}
}

/*
	"date_histogram(field=insert_time,interval=1d)": {
		"date_histogram": {
			"field": "insert_time",
			"interval": "1d",
			"format": "yyyy-MM-dd HH:mm:ss"
		}
	}
*/
func buildGroupByFuncDateHistogram(field, interval, format string, histoAlias string) H {
	m := H{
		"date_histogram": H{
			"field":    field,
			"interval": interval,
			"format":   format,
		},
	}
	return H{histoAlias: m}
}

/*
	"range(age,20,25,30,35,40)": {
		"range": {
			"field": "age",
			"ranges": [
				{
					"from": 20,
					"to": 25
				},
				{
					"from": 25,
					"to": 30
				},
			]
		},
*/
func buildGroupByFuncRange(fieldName string, rangeAggAlias string, rangeArr []int) H {
	m := H{"field": fieldName}
	rangeMapArr := make([]H, len(rangeArr)-1)
	for i := 0; i < len(rangeMapArr); i++ {
		rangeMapArr[i] = H{
			"from": rangeArr[i],
			"to":   rangeArr[i+1],
		}
	}
	m["ranges"] = rangeMapArr
	return H{rangeAggAlias: m}
}

/*
	"date_range(field=insert_time,format=yyyy-MM-dd,2014-08-18,2014-08-17,now-8d,now-7d,now-6d,now)": {
		"date_range": {
			"field": "insert_time",
			"ranges": [
				{
					"from": "2014-08-18",
					"to": "2014-08-17"
				},
				{
					"from": "2014-08-17",
					"to": "now-8d"
				},
			],
			"format": "yyyy-MM-dd"
		}
*/
func buildGroupByFuncDateRange(fieldName string, dateRangeAlias string, format string, dateRangeArr []string) H {
	m := H{
		"date_range": H{
			"field": fieldName,
		},
		"format": "yyyy-MM-dd",
		"ranges": H{},
	}

	rangeMapArr := make([]H, len(dateRangeArr)-1)
	for i := 0; i < len(rangeMapArr); i++ {
		rangeMapArr[i] = H{
			"from": dateRangeArr[i],
			"to":   dateRangeArr[i+1],
		}
	}
	m["ranges"] = rangeMapArr
	if format != "" {
		m["format"] = format
	}

	return H{dateRangeAlias: m}
}

// 可以并列在最内部的aggregation列表中
func buildSelectFuncObj(funcName string, fieldName string, distinct bool) H {
	switch funcName {
	case "count":
		if distinct == true {
			return H{
				"cardinality": H{
					"field": fieldName,
				},
			}
		}
		return H{
			"value_count": H{
				"field": "_index",
			},
		}
	default:
		return H{
			funcName: H{
				"field": fieldName,
			},
		}
	}
}

// var aggMap
// aggMap['aggs'] = aggObj
// aggObj['aggs'] = nestedAggObj

func main() {
	m := buildGroupByField("id", "user_id", 100)
	fmt.Println(m)
	res, _ := json.Marshal(m)
	fmt.Println(string(res))
	m["aggregation"] = buildGroupByField("name", "name", 0)
	res, _ = json.Marshal(m)
	fmt.Println(string(res))
}
