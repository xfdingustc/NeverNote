package controllers
import (
	"github.com/revel/revel"
	"github.com/xfdingustc/NeverNote/app/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/xfdingustc/NeverNote/app/service"
	"github.com/xfdingustc/NeverNote/app/utils"
)

type Notebook struct {
	BaseController
}


func (c Notebook) AddNotebook(parentNotebookId string) revel.Result {
	notebook := models.Notebook{
		UserId: c.GetObjectUserId(),
	}

	if parentNotebookId != "" {
		notebook.ParentNotebookId = bson.ObjectIdHex(parentNotebookId)
	}

	ret, notebook := service.NotebookS.AddNotebook(notebook)

	if ret {
		utils.Log("true")
		return c.RenderJson(notebook)
	} else {
		utils.Log("false")
		return c.RenderJson(false)
	}
}
