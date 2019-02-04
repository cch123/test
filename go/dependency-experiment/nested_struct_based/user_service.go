package main

type userService struct {
	UserID func(phone string) int64
}

type onlineUserService struct{}

func (ous onlineUserService) UserID(phone string) int64 {
	return 1
}

type testUserService struct{}

func (tus testUserService) UserID(phone string) int64 {
	return 2
}
