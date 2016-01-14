package service
import (
	"github.com/xfdingustc/NeverNote/app/utils"
	"github.com/xfdingustc/NeverNote/app/models"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"errors"
)

type AuthService struct {

}

func (this *AuthService) Register(email, pwd, fromUserId string) (bool, string) {
	if userService.IsExistsUser(email) {
		return false, "userHasBeenRegistered-" + email
	}

	encryptedPassword := utils.EncryptPassword(pwd)
	if encryptedPassword == "" {
		return false, "GenerateEncryptedPassword error"
	}

	user := models.User{UserId: bson.NewObjectId(), Email: email, Username: email, Password: encryptedPassword}

	return this.register(user)
}


func (this *AuthService) register(user models.User) (bool, string) {
	if userService.AddUser(user) {
		userId := user.UserId.Hex()
		_ = userId
	}

	return true, ""
}

func (this *AuthService) Login(email, password string) (models.User, error)  {
	trimedEmail := strings.Trim(email, "")

	userInfo := userService.GetUserInfoByName(trimedEmail)

	if userInfo.UserId == "" || !utils.ComparePassword(password, userInfo.Password) {
		return userInfo, errors.New("wrong username or password")
	}

	return userInfo, nil

}
