package controllers

import (

	"github.com/revel/revel"
	"github.com/xfdingustc/NeverNote/app/models"
	"strings"
)


type BaseController struct {
	*revel.Controller
}

func (c BaseController) SetLocale() string {
	locale := string(c.Request.Locale)
	lang := locale

	if  strings.Contains(locale, "-") {
		position := strings.Index(locale, "-")
		lang = locale[0:position]
	}

	// only Simple Chinese & English are supported.
	if lang != "zh" && lang != "en" {
		lang = "en"
	}

	c.RenderArgs["locale"] = lang

	return lang
}


func (c BaseController) RenderResponse(response models.Response) revel.Result {
	oldMsg := response.Msg;

	if response.Msg != "" {
		if strings.Contains(response.Msg, "-") {
			msgAndValues := strings.Split(response.Msg, "-")
			if len(msgAndValues) == 2 {
				response.Msg = c.Message(msgAndValues[0], msgAndValues[1])
			} else {
				others := msgAndValues[0:];
				a := make([]interface{}, len(others))
				for i, v := range others {
					a[i] = v
				}

				response.Msg = c.Message(msgAndValues[0], a...)
			}
		} else {
			response.Msg = c.Message(response.Msg);
		}
	}

	if strings.HasPrefix(response.Msg, "???") {
		response.Msg = oldMsg;
	}

	return c.RenderJson(response);
}

func (c BaseController) GetSession(key string) string{
	v, ok := c.Session[key]
	if !ok {
		return ""
	}

	return v
}

func (c BaseController) SetSession(userInfo models.User) {
	if userInfo.UserId.Hex() != "" {
		c.Session["UserId"] = userInfo.UserId.Hex()
		c.Session["Email"] = userInfo.Email
		c.Session["Username"] = userInfo.Username
	}
}
