package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// Person means Person
type Person struct {
	age  int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	var p Person

	rows, err := db.Query(`select * from abc where id = ?`, 10)
	err2 := errors.Wrap(err, "main")
	//err3 := errors.WithStack(err)

	if err != nil {
		fmt.Printf("%#v\n", errors.Cause(err2))
		fmt.Printf("%+v\n", err2)
		fmt.Printf("%s\n", err2)
		return
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&p.name, &p.age)
	}

}
