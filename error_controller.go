package common

import "github.com/astaxie/beego"

type ErrorController struct {
    beego.Controller
}

func (this *ErrorController) Error400() {
  this.Error(map[string]interface{}{"errorCode":400, "error": "Bad Request"})
}

func (this *ErrorController) Error401() {
  this.Error(map[string]interface{}{"errorCode":401, "error": "Unauthorized"})
}
func (this *ErrorController) Error403() {
  this.Error(map[string]interface{}{"errorCode":403, "error": "Access forbidden"})
}

func (this *ErrorController) Error404() {
  this.Error(map[string]interface{}{"errorCode":404, "error": "Not found"})
}

func (this *ErrorController) Error500() {

  this.Error(map[string]interface{}{"errorCode":500, "error": "Internal server error"})
}

func (this *ErrorController) Error502(){
  this.Error(map[string]interface{}{"errorCode":502, "error": "Bad Gateway"})
}

func (this *ErrorController) Error(default_value map[string]interface{}){
  if this.Data["json"] == nil {
    this.Data["json"] = default_value
  }
  this.TplName = "error.tpl"
  this.ServeJSON()
}
