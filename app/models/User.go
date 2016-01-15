package models
import (
	"gopkg.in/mgo.v2/bson"

	"time"
)


type User struct {
	UserId 		bson.ObjectId 	`bson:"_id,omitempty"`
	Email 		string 			`Email`
	Username	string			`Username`
	Password	string			`bson:"Pwd" json:"-"`
	CreatedTime	time.Time		`CreateTime`
}

type UserAndBlogUrl struct {
	User
	BlogUrl 	string			`BlogUrl`
	PostUrl		string			`PostUrl`
}