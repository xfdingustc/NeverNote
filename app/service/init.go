package service

var userService, UserS *UserService

func InitServices() {
	UserS = &UserService{}
	userService = UserS
}
