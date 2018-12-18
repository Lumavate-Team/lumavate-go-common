package controllers

import ()

type SingleUseTokenStruct struct {
	Token string `json:"token"`
}

type SingleUseTokenController struct {
	LumavateController
}

func (this *SingleUseTokenController) Post() {
	lr := this.GetRequest()
	token_obj, code := lr.GetSingleUseToken()

	if code == 200 {
		this.Ctx.Output.SetStatus(200)
		result := SingleUseTokenStruct{"LUMA-SUT1 " + token_obj.Payload.Data.Token + " UrlRef=" + this.Ctx.Input.Param(":url_ref")}
		this.Data["json"] = &result
		this.ServeJSON()
	} else if code == 401 {
		this.Ctx.Output.SetStatus(401)
		result := SingleUseTokenStruct{""}
		this.Data["json"] = &result
		this.ServeJSON()
	} else if code == 403 {
		this.Ctx.Output.SetStatus(403)
		result := SingleUseTokenStruct{""}
		this.Data["json"] = &result
		this.ServeJSON()
	} else {
		this.Ctx.Output.SetStatus(500)
		result := SingleUseTokenStruct{""}
		this.Data["json"] = &result
		this.ServeJSON()
	}
}
