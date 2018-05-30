package common

import "github.com/astaxie/beego"

type ErrorController struct {
    beego.Controller
}

func (this *ErrorController) Error404() {
  if this.Data["json"] == nil {
    this.Data["json"] = map[string]interface{}{"errorCode": 404, "error": "Not found"}
  }
  this.TplName = "error.tpl"
  this.ServeJSON()
}

func (this *ErrorController) Error403() {

  if this.Data["json"] == nil {
    this.Data["json"] = map[string]interface{}{"errorCode": 403, "error": "Access forbidden"}
  }
    this.TplName = "error.tpl"
    this.ServeJSON()
}

func (this *ErrorController) Error500() {

  if this.Data["json"] == nil {
    this.Data["json"] = map[string]interface{}{"errorCode": 500, "error": "Internal server error"}
  }
  this.TplName = "error.tpl"
  this.ServeJSON()
}
