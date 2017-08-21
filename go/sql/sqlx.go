package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Card struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	PicURL string `db:"pic_url"`
}

func main() {
	var cards []Card
	db, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/card_shop")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}

	//err = db.NamedQuery(&cards, "select * from card where id > :id", map[string]interface{}{"id": "80"})
	m := map[string]interface{}{
		"id":  10,
		"ids": []int{78, 79, 80},
	}

	//可以封装一次，在这里一次搞定
	ft, args, err := sqlx.Named("select * from card where id > :id and id in (:ids)", m)
	fmt.Println(ft, args, err)
	ft, args, err = sqlx.In(ft, args...)

	fmt.Println(ft, args, err)
	err = db.Select(&cards, ft, args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("%v\n", cards[0].Name)
	fmt.Printf("%v\n", cards)
	//println(len(cards))

}
