package main

// dependencies that contains other dependencies
type pushService interface {
	PushMSG(msg string) error
	InitDependency(us userService) error
}

type onlinePushService struct {
	userServiceInstance userService
}

func (ops onlinePushService) PushMSG(msg string) error {
	userID := ops.userServiceInstance.UserID("14012321231")
	println("push to user", userID)
	return nil
}

func (ops onlinePushService) InitDependency(us userService) error {
	ops.userServiceInstance = us
	println("init the user service dependency to the input param")
	return nil
}

type testPushService struct {
	userServiceInstance userService
}

func (tps testPushService) PushMSG(msg string) error {
	userID := tps.userServiceInstance.UserID("14012321231")
	println("push to user", userID)
	return nil
}

func (tps testPushService) InitDependency(us userService) error {
	tps.userServiceInstance = us
	println("init the user service dependency to the input param")
	return nil
}
