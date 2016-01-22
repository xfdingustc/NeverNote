package service
import (
	. "github.com/xfdingustc/NeverNote/app/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/xfdingustc/NeverNote/app/database"
	"time"
	"github.com/xfdingustc/NeverNote/app/utils"
)

type NotebookService struct {

}

func (this *NotebookService) GetNotebooks (userId string) SubNotebooks {
	userNotebooks := []Notebook{}

	database.Notebooks.Find(bson.M{"UserId": bson.ObjectIdHex(userId)}).All(&userNotebooks)

	if len(userNotebooks) == 0 {
		return nil
	}

	return ParseAndSortNotebooks(userNotebooks, true, true)
}

func (this *NotebookService) AddNotebook(notebook Notebook) (bool, Notebook) {
	if notebook.NotebookId == "" {
		notebook.NotebookId = bson.NewObjectId()
	}

	now := time.Now()
	notebook.CreatedTime = now;
	notebook.UpdatedTime = now;

	utils.LogJson(notebook)

	err := database.Notebooks.Insert(notebook)
	if err != nil {
		utils.Log(err)
		return false, notebook
	}
	return true, notebook
}


func ParseAndSortNotebooks(userNotebooks []Notebook, noParentDelete, needSort bool) SubNotebooks {
	userNotebooksMap := make(map[bson.ObjectId]*Notebooks, len(userNotebooks))

	for _, each := range userNotebooks {
		newNotebooks := Notebooks{Subs: SubNotebooks{}}
		newNotebooks.NotebookId = each.NotebookId
		newNotebooks.Title = each.Title
		newNotebooks.UserId = each.UserId
		newNotebooks.ParentNotebookId = each.ParentNotebookId

		userNotebooksMap[each.NotebookId] = &newNotebooks
	}

//	for id, each := range userNotebooksMap {
//		if each.ParentNotebookId.Hex() != "" {
//			if
//		}
//	}

	final := make(SubNotebooks, len(userNotebooksMap))
	i := 0
	for _, each := range userNotebooksMap {
		final[i] = each
		i++
	}

	return final
}
