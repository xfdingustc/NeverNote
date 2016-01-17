package service
import (
	"github.com/xfdingustc/NeverNote/app/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/xfdingustc/NeverNote/app/database"
)

type NotebookService struct {

}

func (this *NotebookService) GetNotebooks (userId string) models.SubNotebooks {
	userNotebooks := []models.Notebook{}

	database.Notebooks.Find(bson.M{"UserId": bson.ObjectIdHex(userId)}).All(&userNotebooks)

	if len(userNotebooks) == 0 {
		return nil
	}

	return ParseAndSortNotebooks(userNotebooks, true, true)
}


func ParseAndSortNotebooks(userNotebooks []models.Notebook, noParentDelete, needSort bool) models.SubNotebooks {
	userNotebooksMap := make(map[bson.ObjectId]*models.Notebooks, len(userNotebooks))

	for _, each := range userNotebooks {
		newNotebooks := models.Notebooks{Subs: models.SubNotebooks{}}
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

	return nil
}
