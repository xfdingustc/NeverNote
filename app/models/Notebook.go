package models
import (
	"gopkg.in/mgo.v2/bson"
	"time"
)


type Notebook struct {
	NotebookId			bson.ObjectId		`bson:"_id,omitempty"`
	ParentNotebookId	bson.ObjectId		`bson:"ParentNotebookId,omitempty"`
	UserId 				bson.ObjectId		`bson:"UserId"`

	Title				string				`Title`

	CreatedTime			time.Time			`CreatedTime,omitempty`
	UpdatedTime			time.Time			`UpdatedTime,omitempty`

}


type SubNotebooks []*Notebooks

type Notebooks struct {
	Notebook
	Subs SubNotebooks
}