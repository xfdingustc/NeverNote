package service
import (
	"github.com/xfdingustc/NeverNote/app/utils"
	"github.com/xfdingustc/NeverNote/app/models"
	"gopkg.in/mgo.v2/bson"
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
