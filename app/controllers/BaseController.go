package controllers

import (

	"github.com/revel/revel"
	"github.com/xfdingustc/NeverNote/app/info"
	"strings"
)


type BaseController struct {
	*revel.Controller
}


func (c BaseController) RenderResponse(response info.Response) revel.Result {
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
