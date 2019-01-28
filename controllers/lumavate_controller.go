package controllers

import (
	_ "errors"
	"fmt"
	"github.com/Lumavate-Team/lumavate-go-common/api_core"
	"github.com/astaxie/beego"
	"os"
	"strings"
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
	lr := this.GetRequest()
	use_single_token := lr.ExtractSingleTokenFlag(single_token)
	return lr.Get(url, use_single_token)
}

func (this *LumavateController) LumavatePost(url string, payload []byte, single_token ...bool) ([]byte, string) {
	lr := this.GetRequest()
	use_single_token := lr.ExtractSingleTokenFlag(single_token)
	return lr.Post(url, payload, use_single_token)
}

func (this *LumavateController) LumavatePut(url string, payload []byte, single_token ...bool) ([]byte, string) {
	lr := this.GetRequest()
	use_single_token := lr.ExtractSingleTokenFlag(single_token)
	return lr.Put(url, payload, use_single_token)
}

func (this *LumavateController) LumavateDelete(url string, payload []byte, single_token ...bool) ([]byte, string) {
	lr := this.GetRequest()
	use_single_token := lr.ExtractSingleTokenFlag(single_token)
	return lr.Delete(url, payload, use_single_token)
}

func (this *LumavateController) LumavatePatch(url string, payload []byte, single_token ...bool) ([]byte, string) {
	lr := this.GetRequest()
	use_single_token := lr.ExtractSingleTokenFlag(single_token)
	return lr.Patch(url, payload, use_single_token)
}

func (this *LumavateController) LumavateGetData() []byte {
	this.LumavateInit()
	data, status := this.LumavateGet(this.GetWidgetDataUrl())
	switch status {
	case "200":
		return data
	case "401":
		fmt.Println("401")
		this.NoAuthRedirect()
	default:
		this.Abort(status)
	}
	return []byte{}
}

// returns status codes instead of eating codes and doing redirects
// useful for when a widget is calling to an api widget like in data-widget
func (this *LumavateController) LumavateApiGetData() ([]byte, string) {
	this.LumavateInit()
	data, status := this.LumavateGet(this.GetWidgetDataUrl())
	switch status {
	case "200":
		return data, "200"
	case "401":
		return []byte{}, "401"
	default:
		return []byte{}, status
	}
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

// attempt to get token from header first.
// if it doesn't exist their then fallback to cookie
func (this *LumavateController) GetRequest() api_core.LumavateRequest {
	auth_header := this.Ctx.Input.Header("Authorization")

	if len(auth_header) > 0 && strings.HasPrefix(auth_header, "Bearer ") {
		return api_core.LumavateRequest{strings.TrimPrefix(auth_header, "Bearer "), this.Ctx.Input.Host()}
	}

	return api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt"), this.Ctx.Input.Host()}
}

func (this *LumavateController) MustHaveValidSingleUseToken() {
	token := this.Ctx.Input.Header("Experience-Token")
	if token == "" {
		this.Abort("403")
	}

	fmt.Println("HOST: " + this.Ctx.Input.Host())
	lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt"), this.Ctx.Input.Host()}
	_, status := lr.Get(fmt.Sprintf("/pwa/v1/single-use-token/%v", token), false)
	if status == "400" {
		this.Data["json"] = map[string]interface{}{"errorCode": 403, "error": "Invalid Token"}
		this.Abort("403")
	}
}