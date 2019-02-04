package main

import "fmt"

// problem: hard to express nested dependencies
// if dependency has dependency
// you'll have to nest some struct
type dependencyList struct {
	GetUserID    func(userID int) int64
	GetWhiteList func(date string) map[int]int64
	GetLimited   func(driverID int64) bool
	PushMSG      func(userID int64) error
}

var exactDependency dependencyList
var env = "test"

func main() {
	switch env {
	case "test":
		testInitDependency()
	case "prod":
		onlineInitDependency()
	}
	m := exactDependency.GetWhiteList("abcabc")
	fmt.Println(m)
}
