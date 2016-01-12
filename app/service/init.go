package service

var userService, UserS *UserService
var authService, AuthS *AuthService

func InitServices() {
	UserS = &UserService{}
	AuthS = &AuthService{}
	userService = UserS
	authService = AuthS
}
