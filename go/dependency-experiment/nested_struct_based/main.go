package main

type dependencyList struct {
	orderServiceInstance orderService
	pushServiceInstance  pushService
}

// dependency list manager
var exactDependency dependencyList

var env = "test"

func main() {
	switch env {
	case "prod":
		normalInit()
	case "test":
		testInit()
	}
}
