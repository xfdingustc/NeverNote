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
		return c.RenderJson(notebook)
	} else {
		utils.Log("false")
		return c.RenderJson(false)
	}
}

func (c Notebook) DeleteNotebook(notebookId string) revel.Result {
	ok, msg := service.NotebookS.DeleteNotebook(c.GetUserId(), notebookId);
	utils.Log(ok)
	utils.Log(msg)
	return c.RenderJson(models.Response{Ok: ok, Msg: msg})
}

func (c Notebook) UpdateNotebookTitle(notebookId, title string) revel.Result {
	userId := c.GetUserId()
	return c.RenderJson(service.NotebookS.UpdateNotebookTitle(notebookId, userId, title));
}
