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
	ft, args, err := sqlx.Named("select * from card where id > :id", map[string]interface{}{"id": 1})
	//fmt.Println(fmt, args, err)
	fmt.Println(ft, args[0].(int), err)
	err = db.Select(&cards, ft, args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("%v\n", cards[0].Name)
	fmt.Printf("%v\n", cards)
	//println(len(cards))

}
