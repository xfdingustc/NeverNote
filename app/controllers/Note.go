package controllers
import (
	"github.com/revel/revel"

	. "github.com/xfdingustc/NeverNote/app/service"
	"github.com/xfdingustc/NeverNote/app/utils"
)

type Note struct {
	BaseController
}



func (c Note) Index(noteId, online string) revel.Result {
	c.SetLocale()

	//userInfo := c.GetUserAndBlogUrl()
	user := c.GetUser()
	utils.LogJson(user)
	//userId := userInfo
	userId := user.UserId.Hex();
	if userId == "" {
		return c.Redirect("/login")
	}

	notebooks := NotebookS.GetNotebooks(userId)


	c.RenderArgs["notebooks"] = notebooks

	return c.RenderTemplate("note/note.html")
}
