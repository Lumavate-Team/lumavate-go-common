package common

import "github.com/astaxie/beego"

type ErrorController struct {
    beego.Controller
}

func (this *ErrorController) Error404() {
    this.Data["json"] = map[string]interface{}{"errorCode": 404, "error": "Not Found"}
    this.ServeJSON()
}

func (this *ErrorController) Error403() {
    this.Data["json"] = map[string]interface{}{"errorCode":403, "error":"Access forbidden"}
    this.ServeJSON()
}
