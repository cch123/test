package main

//godoc encoding/json Marshal
//返回的是byte数组，需要手动转成string
import "encoding/json"

import "fmt"

type Server struct {
	ID         int    `json:"-"`
	ServerName string `json:"serverName"`
	ServerIP   string `json:"ip"`
}

func main() {
	s := Server{
		ID:         1,
		ServerName: `1.0 "server"`,
		ServerIP:   `127.0.0.1`,
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
}
