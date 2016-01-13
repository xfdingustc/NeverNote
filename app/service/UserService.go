package service
import (
	"strings"
	_ "github.com/xfdingustc/NeverNote/app/models"
)

type UserService struct {

}




func (this *UserService) GetUserId(email string) string {
	email = strings.ToLower(email)
//	user := models.User{}
//	db.GetByQ(db.Users, bson.M{"Email": email}, &user)
//	return user.UserId.Hex()
	return ""
}


func (this *UserService) IsExistsUser(email string) bool {
	if this.GetUserId(email) == "" {
		return false
	}
	return true
}

