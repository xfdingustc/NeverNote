package controllers

import (
	"github.com/revel/revel"
	//"github.com/xfdingustc/NeverNote/app/info"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) DoRegister(email, password string) revel.Result {
	//re := info.NewRe()

//	if re.Ok, re.Msg = Vd("email", email); !re.OK {
//		return c.RenderRe(re);
//	}
	return  nil;
}

func (c Auth) Register(from, iu string) revel.Result {
	//c.SetLocale();
	return c.RenderTemplate("home/register.html");
}
