package controllers

import "github.com/revel/revel"

type Index struct {
	BaseController
}


func (c Index)Index() revel.Result  {
	return c.RenderTemplate("home/index.html");
}