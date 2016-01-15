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

func (this *UserService) GetUserInfoByName(email string) models.User {
	emailOrUsername := strings.ToLower(email)

	user := models.User{}
	if strings.Contains(email, "@") {
		database.GetByQ(database.Users, bson.M{"Email": emailOrUsername}, &user)
	} else {
		database.GetByQ(database.Users, bson.M{"Username": emailOrUsername}, &user)
	}

	return user
}

func (this *UserService) GetUserInfo(userId string) models.User {
	user := models.User{}
	database.Get(database.Users, userId, &user)
	return user
}

//func (this *UserService) GetUserAndBlogUrl(userId string) models.UserAndBlogUrl {
//	user := this.GetUserInfo(userId)
//
//	userBlog := blogService.GetUserBlog(userId)
//
//}

