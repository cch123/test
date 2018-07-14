package xxx

import "fmt"

var (
	getUserFunc     func() bool
	validateFunc    func() bool
	createOrderFunc func() bool
)

// Process ...
func Process() {
	fmt.Println(getUserFunc())
	fmt.Println(validateFunc())
	fmt.Println(createOrderFunc())
}
