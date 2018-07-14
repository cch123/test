// +build test

package xxx

func init() {
	getUserFunc = func() bool { return false }
	validateFunc = func() bool { return false }
	createOrderFunc = func() bool { return false }
}
