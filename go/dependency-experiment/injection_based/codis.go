package main

import "fmt"

type codisService struct {
	connOptions []string
}

func initCodis() (*codisService, error) {
	fmt.Println("init codis")
	return &codisService{
		connOptions: []string{"yyy", "zzz"},
	}, nil
}
