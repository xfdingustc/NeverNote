package controllers
import (
	"github.com/revel/revel"
	"github.com/xfdingustc/NeverNote/app/service"
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

	notebooks := service.NotebookS.GetNotebooks(userId)

	return c.RenderTemplate("note/note.html")
}
