package controllers
import (
	"github.com/revel/revel"
)

type Note struct {
	BaseController
}



func (c Note) Index(noteId, online string) revel.Result {
	c.SetLocale()

	//userInfo := c.GetUserAndBlogUrl()
	user := c.GetUser()
	//userId := userInfo
	userId := user.UserId.Hex();
	if userId == "" {
		return c.Redirect("/login")
	}

//	notebooks := service.NotebookS.GetNotebooks(userId)

//	c.RenderArgs["notebooks"] = notebooks

	return c.RenderTemplate("note/note.html")
}
