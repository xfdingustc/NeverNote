package service

var userService, UserS *UserService
var authService, AuthS *AuthService
var blogService, BlogS *BlogService
var notebookService, NotebookS *NotebookService

func InitServices() {
	UserS = &UserService{}
	AuthS = &AuthService{}
	BlogS = &BlogService{}
	NotebookS = &NotebookService{}
	userService = UserS
	authService = AuthS
	blogService = BlogS
	notebookService = NotebookS
}
