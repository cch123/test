package main

import "menteslibres.net/gosexy/redis"
import "log"

var client *redis.Client
var host = "127.0.0.1"
var port = uint(6379)
var s string

func main() {
	client = redis.New()
	defer client.Quit()
	err := client.Connect(host, port)
	if err != nil {
		println(err.Error())
		return
	}
	println("connect to redis server")
	//example ping
	println("sending ping")
	s, err = client.Ping()
	if err != nil {
		println("Ping error")
		return
	}
	println("received " + s)
	//end of example ping

	//example set get
	println("set Xargin")
	client.Set("Xargin", 1)
	s, err = client.Get("Xargin")
	if err != nil {
		println("get Xargin error")
		return
	}
	println("get Xargin = " + s)
	//end of exmplae set get

	//example incr
	println("increment Xargin")
	client.Incr("Xargin")
	s, err = client.Get("Xargin")
	if err != nil {
		println("get Xargin err")
		return
	}
	println("get Xargin = " + s)
	//end
	log.Println("end")
}
