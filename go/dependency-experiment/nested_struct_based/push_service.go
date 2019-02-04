package main

// dependencies that contains other dependencies
type pushService struct {
	PushMSG        func(msg string) error
	InitDependency func(us userService) error
}

func onlinePushMSG(msg string) error {
	println("push to user")
	return nil
}

func onlineInitDependency(us userService) error {
	println("init the user service dependency to the input param")
	return nil
}

type testPushService struct {
	userServiceInstance userService
}

func testPushMSG(msg string) error {
	println("push to user")
	return nil
}

func testInitDependency(us userService) error {
	println("init the user service dependency to the input param")
	return nil
}
