package service
import (
	"strings"
	"github.com/xfdingustc/NeverNote/app/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/xfdingustc/NeverNote/app/database"
	"time"
)

type UserService struct {

}

func (this *UserService) AddUser(user models.User) bool {
	if user.UserId == "" {
		user.UserId = bson.NewObjectId()
	}

	user.CreatedTime = time.Now()

	if user.Email != "" {
		user.Email = strings.ToLower(user.Email)


	}

	return database.Insert(database.Users, user)
}


func (this *UserService) GetUserId(email string) string {
	email = strings.ToLower(email)
	user := models.User{}
	database.GetByQ(database.Users, bson.M{"Email": email}, &user)
	return user.UserId.Hex()

}


func (this *UserService) IsExistsUser(email string) bool {
	if this.GetUserId(email) == "" {
		return false
	}
	return true
}

