package main

import (
	"database/sql"
	"fmt"

	"go.uber.org/dig"
)

type rateDBParam struct {
	dig.In
	DBInstance *sql.DB `name:"rate_db"`
}

type rateService struct {
	DBInstance *sql.DB
}

func initRateService(cs *codisService, db rateDBParam) (*rateService, error) {
	fmt.Println("init rate service")
	return &rateService{DBInstance: db.DBInstance}, nil
}

func initRateDatabase() (*sql.DB, error) {
	fmt.Println("init rate database")
	return &sql.DB{}, nil
}
