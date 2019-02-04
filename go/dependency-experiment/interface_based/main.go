package main

type dependencyList struct {
	priceServiceInstance priceService
	pushServiceInstance  pushService
	orderServiceInstance orderService
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
	exactDependency.priceServiceInstance.GetPrice(1)
}
