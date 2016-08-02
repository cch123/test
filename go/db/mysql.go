package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "aaaa:aaaa@/houchuwang")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("select * from user limit 1")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	rows, err := stmt.Exec()
	println(rows)
}
