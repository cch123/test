package main

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

// validate 只有在 validate的内容是string时才会有效
type collectStarParams struct {
	BizType string `valid:"numeric,required"`
}

func main() {
	var cp = collectStarParams{}
	ok, err := valid.ValidateStruct(cp)
	fmt.Println(ok, err)
}
