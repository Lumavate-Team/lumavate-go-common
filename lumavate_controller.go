package common

import (
  "github.com/astaxie/beego"
  "fmt"
  "os"
  "github.com/Lumavate-Team/lumavate-go-common/api_core"
)

type LumavateController struct {
  beego.Controller
}

func (this *LumavateController) LumavateInit() {
	this.Data["WidgetUrlPrefix"] = os.Getenv("WIDGET_URL_PREFIX")

  this.LayoutSections = make(map[string]string)
  this.LayoutSections["HtmlHead"] = ""
  this.LayoutSections["HeaderContent"] = ""
  this.LayoutSections["FooterContent"] = ""
  this.LayoutSections["Scripts"] = ""
}

func (this *LumavateController) LumavateGet(url string) ([]byte, string) {
  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt")}
  return lr.Get(url)
}

func (this *LumavateController) LumavatePost(url string, payload []byte) ([]byte, string) {
  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt")}
  return lr.Post(url, payload)
}

func (this *LumavateController) LumavatePut(url string, payload []byte) ([]byte, string) {
  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt")}
  return lr.Put(url, payload)
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
