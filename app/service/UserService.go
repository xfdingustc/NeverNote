package service
import (
	"strings"
)

type UserService struct {

}




func (this *UserService) GetUserId(email string) string {
	email = strings.ToLower(email)
//	user := info.User{}
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

