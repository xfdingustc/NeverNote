package controllers
import "github.com/revel/revel"

type Note struct {
	BaseController
}



func (c Note) Index(noteId, online string) revel.Result {
	c.SetLocale()

	//userInfo =

	return c.RenderTemplate("note/note.html")
}
