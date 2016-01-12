package service

type AuthService struct {

}

func (this *AuthService) Register(email, pwd, fromUserId string) (bool, string) {
	if userService.IsExistsUser(email) {
		return false, "userHasBeenRegistered-" + email
	}

	return true, ""
}
