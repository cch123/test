package main

import (
	"database/sql"
	"fmt"

	"go.uber.org/dig"
)

func initStarDatabase() (*sql.DB, error) {
	fmt.Println("init star database")
	return &sql.DB{}, nil
}

type starDBParam struct {
	dig.In
	DBInstance *sql.DB `name:"star_db"`
}

type starService struct {
	CodisInstance *codisService
	DBInstance    *sql.DB
}

var ss starService

func initStarService(cs *codisService, db starDBParam) (*starService, error) {
	fmt.Println("init star service")
	ss = starService{
		CodisInstance: cs,
		DBInstance:    db.DBInstance,
	}

	return &starService{
		CodisInstance: cs,
		DBInstance:    db.DBInstance,
	}, nil
}
