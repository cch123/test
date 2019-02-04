package main

type userService interface {
	UserID(phone string) int64
}

type onlineUserService struct{}

func (ous onlineUserService) UserID(phone string) int64 {
	return 1
}

type testUserService struct{}

func (tus testUserService) UserID(phone string) int64 {
	return 2
}
