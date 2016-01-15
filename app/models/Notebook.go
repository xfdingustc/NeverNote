package models
import (
	"gopkg.in/mgo.v2/bson"
	"time"
)


type Notebook struct {
	NotebookId		bson.ObjectId		`bson:"_id,omitempty"`
	UserId 			bson.ObjectId		`bson:"UserId"`
	CreatedTime		time.Time			`CreatedTime,omitempty`
	UpdatedTime		time.Time			`UpdatedTime,omitempty`

}
