package common

import (
  "github.com/astaxie/beego"
  "fmt"
 _ "errors"
  "os"
  "github.com/Lumavate-Team/lumavate-go-common/api_core"
)

type LumavateController struct {
  beego.Controller
}

func (this *LumavateController) LumavateInit() {
	this.Data["WidgetUrlPrefix"] = os.Getenv("WIDGET_URL_PREFIX")
  this.Data["CacheKey"] = os.Getenv("PUBLIC_KEY")

  this.LayoutSections = make(map[string]string)
  this.LayoutSections["HtmlHead"] = ""
  this.LayoutSections["HeaderContent"] = ""
  this.LayoutSections["FooterContent"] = ""
  this.LayoutSections["Scripts"] = ""
}

func (this *LumavateController) LumavateGet(url string, single_token ...bool) ([]byte, string) {
  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt")}

  use_single_token := lr.ExtractSingleTokenFlag(single_token)
  return lr.Get(url, use_single_token)
}

func (this *LumavateController) LumavatePost(url string, payload []byte, single_token ...bool) ([]byte, string) {
  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt")}

  use_single_token := lr.ExtractSingleTokenFlag(single_token)
  return lr.Post(url, payload, use_single_token)
}

func (this *LumavateController) LumavatePut(url string, payload []byte, single_token ...bool) ([]byte, string) {
  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt")}

  use_single_token := lr.ExtractSingleTokenFlag(single_token)
  return lr.Put(url, payload, use_single_token)
}

func (this *LumavateController) LumavateDelete(url string, payload []byte, single_token ...bool) ([]byte, string) {
  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt")}

  use_single_token := lr.ExtractSingleTokenFlag(single_token)
  return lr.Delete(url, payload, use_single_token)
}

func (this *LumavateController) LumavateGetData() []byte {
  this.LumavateInit()
  data, status := this.LumavateGet(this.GetWidgetDataUrl())
  switch status {
  case "200":
    return data
  case "401":
    this.NoAuthRedirect()
  default:
    this.Abort(status)
  }
  return []byte{}
}

func (this *LumavateController) GetSelfUrl() string {
	return this.Ctx.Input.URI()
}

func (this *LumavateController) GetNoAuthRedirectUrl() string {
  return fmt.Sprintf("%s%s?u=%s",
    os.Getenv("PROTO"),
    this.Ctx.Input.Host(),
    this.GetSelfUrl(),
    )
}

func (this *LumavateController) NoAuthRedirect() {
  this.Ctx.Redirect(302, this.GetNoAuthRedirectUrl())
}

func (this *LumavateController) GetWidgetDataUrl() string {
  return fmt.Sprintf("/pwa/v1/widget-instances/%s",
    this.Ctx.Input.Param(":wid"),
  )
}

func (this *LumavateController) MustHaveValidSingleUseToken() {
  token := this.Ctx.Input.Header("Experience-Token")
  if token == "" {
    this.Abort("403")
  }
  
  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt")}
  _, status := lr.Get(fmt.Sprintf("/pwa/v1/single-use-token/%v", token), false)
  fmt.Println("STATUS ", status)
  if status == "400" {
//    this.Abort("403")
    this.Ctx.ResponseWriter.WriteHeader(403)
    this.Data["json"] = map[string]interface{}{"Error":"Invalid token"}
	  this.ServeJSON()
  }
}


