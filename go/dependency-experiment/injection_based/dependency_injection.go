package main

import (
	"fmt"

	"go.uber.org/dig"
)

func init() {
	c := dig.New()
	if err := c.Provide(initCodis); err != nil {
		fmt.Println(err, "init codis service failed")
	}

	c.Provide(initRateDatabase, dig.Name("star_db"))
	c.Provide(initStarDatabase, dig.Name("rate_db"))
	c.Provide(initStarService)
	c.Provide(initRateService)

	err := c.Invoke(initStarService)
	err = c.Invoke(initRateService)

	if err != nil {
		fmt.Println(err, "invoke rate service failed")
	}

	// fmt.Println(ss.CodisInstance.connOptions)
	//fmt.Println(c)
	fmt.Println(ss.DBInstance)

}

func main() {
}
