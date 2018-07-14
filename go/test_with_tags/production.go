// +build !test

package xxx

func init() {
	getUserFunc = func() bool { return true }
	validateFunc = func() bool { return true }
	createOrderFunc = func() bool { return true }
}
