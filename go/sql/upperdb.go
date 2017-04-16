package main

import (
	"fmt"
	"os"

	"reflect"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var settings = mysql.ConnectionURL{
	Host:     "localhost",
	Database: "card_shop",
	User:     "root",
	Password: "",
}

// Card is
type Card struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// v3直接使用了context，所以应该是1.7以上才成
// v3直接使用了context，所以应该是1.7以上才成
// v3直接使用了context，所以应该是1.7以上才成
// v3直接使用了context，所以应该是1.7以上才成
func main() {
	// sess is cocurrent safe
	// 所以项目全局只需要有一个session
	sess, err := mysql.Open(settings)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 连接设置
	sess.SetMaxOpenConns(1000)
	sess.SetMaxIdleConns(999)

	// 打印sql查询日志，调试的时候非常实用
	sess.SetLogging(true)
	defer sess.Close()

	// 支持直接和struct或者map绑定
	var cards []Card
	var cardMapList []map[string]interface{}
	var c Card
	// 简单查询
	//err = sess.Collection("card").Find().All(&cards)
	//q := sess.Select("id", "name").From("card").Where(db.Cond{"id > ": 10, "name > ": "小"})
	q := sess.Select("id", "name").From("card").And("id > ?", 10).And("name > ?", "小")
	//q = sess.Select("id", "name").From("card").Where("id > ?", 10)
	//q = sess.Select("id", "name").From("card").Where("id > ?", 10)
	// <=> 是mysql中null安全的=操作方言
	//q = sess.Select("id", "name").From("card").Where("id <=> ?", nil)
	q = sess.Select("id", "name").From("card").Where("id in", []int{70, 100, 200})
	q = sess.Select("id", "name").From("card").Where("id<=> ?", 70)
	//q = sess.Select("id", "name").From("card").Where("id > ? and name > ?", 10, "小")
	// 反操作是 NOT( xxx <=> yyy)

	// 复合查询
	// 具体可以查看db.And，db.Or的定义
	// 不要全是where，因为where会进行覆盖，相关的讨论：
	// https://github.com/upper/db/issues/357
	q = sess.Select("id", "name").From("card").
		Where(
		db.And(
			db.Cond{"id > ": 0},
			db.Or(
				db.Cond{"name > ": "小"},
				db.Cond{"name <": "白"},
			),
		))
	q.All(&cards)
	q.One(&c)
	q.All(&cardMapList)

	fmt.Printf("%#v\n", cards)
	fmt.Printf("%#v\n", c)

	//绑定到map的时候，注意处理 interface{}到字符串还会比较蛋疼，需要自己转换为string
	for _, v := range cardMapList {
		for mapKey, mapVal := range v {
			if reflect.TypeOf(mapVal).Kind() == reflect.Slice &&
				reflect.TypeOf(mapVal).Elem().Kind() == reflect.Uint8 {
				str := string(mapVal.([]byte))
				fmt.Print(mapKey, " : ", str, " ;")
				//fmt.Printf("%v", mapVal)
				//不做转换的话，会以[]uint8数组的形式呈现
			} else {
				fmt.Print(mapKey, " : ", mapVal, ";")
			}
		}
		fmt.Println()
	}

	// INSERT部分
	toInsert := Card{Name: "testests"}
	qq := sess.InsertInto("card").Values(toInsert)
	res, err := qq.Exec()
	fmt.Println(res, err)
	lastInsertID, err := res.LastInsertId()
	fmt.Println("last insert id:", lastInsertID)
	// batch insert
	toInsert2 := []Card{Card{Name: " ttttt"}, Card{Name: "zzzzz"}}
	qqq := sess.InsertInto("card").Values(toInsert2[0]).Values(toInsert2[1])
	res, err = qqq.Exec()
	fmt.Println(res, err)
	// 还有一个batch方法，懒得写了
	// sess.InsertInto("card").Batch()
	// https://github.com/upper/db/issues/190

	// update
	q3 := sess.Update("card").Set(map[string]interface{}{"name": "zzzzz"}).Where("id >", "200")
	res, err = q3.Exec()
	affected, err := res.RowsAffected()
	fmt.Println("affected", affected)
	// update 形如 column = column+1 的 update
	q3 = sess.Update("card").Set("id = id + 100000").Where("id > ", 200)
	res, err = q3.Exec()
	affected, err = res.RowsAffected()

	//if的判断也是支持的
	q3 = sess.Update("card").Set("id = IF(id - 200000 > 0 , id - 200000, id*2)").Where("id >", 200)
	res, err = q3.Exec()
	affected, err = res.RowsAffected()

	// delete
	// delete在业务框架里应该封装一次，避免没有条件的整表delete~
	q4 := sess.DeleteFrom("card").Where("id > ?", 200)
	res, err = q4.Exec()
	affected, err = res.RowsAffected()
	fmt.Println("affected", affected)

}
