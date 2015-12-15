package controllers

import (
	"github.com/revel/revel"
	"github.com/xfdingustc/NeverNote/app/info"
	. "github.com/xfdingustc/NeverNote/app/utils"
	"strings"
)

type Auth struct {
	BaseController
}

func (c Auth) DoRegister(email, password string) revel.Result {
	response := info.NewResponse()

	if response.Ok, response.Msg = Validate("email", email); !response.Ok {
		return c.RenderResponse(response)
	}

	if response.Ok, response.Msg = Validate("password", password); !response.Ok {
		return c.RenderResponse(response)
	}

	email = strings.ToLower(email)

	//response.Ok, response.Msg := au

	return c.RenderResponse(response)
}

func (c Auth) Register(from, iu string) revel.Result {
	//c.SetLocale();
	return c.RenderTemplate("home/register.html");
}
