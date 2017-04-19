// 这个包几乎没什么用，大多数参数的问题是需要在绑定之前就做检查的
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
