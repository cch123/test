package main

import (
	"fmt"
	"os"

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
func main() {
	// sess is cocurrent safe
	sess, err := mysql.Open(settings)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sess.SetMaxOpenConns(1000)
	sess.SetMaxIdleConns(999)
	sess.SetLogging(true)
	defer sess.Close()

	var cards []Card
	var c Card
	//err = sess.Collection("card").Find().All(&cards)
	//q := sess.Select("id", "name").From("card").Where(db.Cond{"id > ": 10, "name > ": "小"})
	q := sess.Select("id", "name").From("card").And("id > ?", 10).And("name > ?", "小")
	//q = sess.Select("id", "name").From("card").Where("id > ?", 10)
	//q = sess.Select("id", "name").From("card").Where("id > ?", 10)
	// <=> 是mysql中null安全的=操作方言
	//q = sess.Select("id", "name").From("card").Where("id <=> ?", nil)
	q = sess.Select("id", "name").From("card").Where("id<=> ?", 70)
	//q = sess.Select("id", "name").From("card").Where("id > ? and name > ?", 10, "小")
	// 反操作是 NOT( xxx <=> yyy)
	q.All(&cards)
	q.One(&c)

	fmt.Printf("%#v\n", cards)
	fmt.Printf("%#v\n", c)
}
