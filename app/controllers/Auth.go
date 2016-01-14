package controllers

import (
	"github.com/revel/revel"
	"github.com/xfdingustc/NeverNote/app/models"
	"github.com/xfdingustc/NeverNote/app/utils"
	"github.com/xfdingustc/NeverNote/app/service"
	"strings"
)

type Auth struct {
	BaseController
}

func (c Auth) DoRegister(email, password, iu string) revel.Result {
	response := models.NewResponse()

	if response.Ok, response.Msg = utils.Validate("email", email); !response.Ok {
		return c.RenderResponse(response)
	}

	if response.Ok, response.Msg = utils.Validate("password", password); !response.Ok {
		return c.RenderResponse(response)
	}

	email = strings.ToLower(email)

	response.Ok, response.Msg = service.AuthS.Register(email, password, iu)

	return c.RenderResponse(response)
}



func (c Auth) Register(from, iu string) revel.Result {
	c.SetLocale();
	c.RenderArgs["from"] = from
	c.RenderArgs["iu"] = iu

	c.RenderArgs["title"] = c.Message("register")
	c.RenderArgs["subTitle"] = c.Message("register");
	return c.RenderTemplate("home/register.html");
}


func (c Auth) Login(email, from string) revel.Result {
	return c.RenderTemplate("home/login.html")
}
