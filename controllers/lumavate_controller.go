package controllers

import (
  _ "errors"
  "fmt"
  "github.com/Lumavate-Team/lumavate-go-common/api_core"
  "github.com/astaxie/beego"
  "os"
  "strings"
  "github.com/Lumavate-Team/lumavate-go-common/models"
  "github.com/Lumavate-Team/lumavate-go-common/properties"
  "encoding/json"
  b64 "encoding/base64"
)

/*
* Internal Types used for managing dynamic components
*/

type DynamicComponent struct {
  Icon       string
  Label      string
  Type       string
  Tags       []string
  Template   string
  Section    string
  Properties []properties.PropertyType
}

type ComponentSetRequest struct {
  Payload struct {
    Data []struct {
      CurrentVersion struct {
        DirectIncludes    []string
        DirectCssIncludes []string
        Distribution      string
        Components        []*DynamicComponent
      }
    }
  }
}

type LumavateController struct {
  beego.Controller
  Components []*DynamicComponent
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

func (this *LumavateController) ParseToken() models.TokenDataStruct {
    token := this.Ctx.GetCookie("pwa_jwt")
		if token == "" {
			token = this.Ctx.Input.Header("Authorization")
			if strings.HasPrefix(token, "Bearer "){
				token = strings.Replace(token, "Bearer ", "", -1)
			}
		}
    token = strings.Split(token, ".")[1]

    // add padding to jwt if number of bytes is not correct
    if i := len(token) % 4; i != 0 {
      token += strings.Repeat("=", 4-i)
    }

    token_data := models.TokenDataStruct{}
    // decode the token and ummarshal into auth struct
    decodedToken, _ := b64.StdEncoding.DecodeString(token)
    if err := json.Unmarshal(decodedToken, &token_data); err != nil {
      fmt.Println(err)
      panic(err)
    }

    return token_data
  
}

func (this *LumavateController) LumavateGetData() models.WidgetData {
  this.LumavateInit()
  data, status := this.LumavateGet(this.GetWidgetDataUrl())
  switch status {
		case "200":
			response := models.WidgetData {}
			json.Unmarshal(data, &response)
			response.Payload.Data.TokenData = this.ParseToken()
			return response
		case "401":
			this.NoAuthRedirect()
		default:
			this.Abort(status)
  }
  return models.WidgetData {}
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
  url := this.Ctx.Input.URI()
  parts := strings.Split(url, "?")
  query := ""
  if len(parts) > 1 {
    query = parts[1]
    return fmt.Sprintf("/pwa/v1/widget-instance-data/%s?%s",
      this.Ctx.Input.Param(":wid"),
      query,
    )
  }

  return fmt.Sprintf("/pwa/v1/widget-instance-data/%s",
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

  lr := api_core.LumavateRequest{this.Ctx.GetCookie("pwa_jwt"), this.Ctx.Input.Host()}
  _, status := lr.Get(fmt.Sprintf("/pwa/v1/single-use-token/%v", token), false)
  if status == "400" {
    this.Data["json"] = map[string]interface{}{"errorCode": 403, "error": "Invalid Token"}
    this.Abort("403")
  }
}

func (this *LumavateController) GetComponentsWithTag(tag string) []*properties.Component {

  components := []*properties.Component{}

  for _, component := range this.Components {
    for _, t := range component.Tags {
      if t == tag {
        components = append(components, &properties.Component{tag, component.Section, component.Type, "", component.Icon, component.Label, component.Properties, ""})
      }
    }
  }

  return components
}

func (this *LumavateController) GetDynamicComponentProperty(tag, name, classification, section, label, help string) *properties.PropertyComponent {

  components := this.GetComponentsWithTag(tag)

  if len(components) == 0 {
    return &properties.PropertyComponent{
      &properties.PropertyBase{tag, classification, section, label, help, ""},
      &properties.Component{}, &properties.PropertyOptionsComponent{[]string{}, []*properties.Component{}, false},
    }
  }

  return &properties.PropertyComponent{
    &properties.PropertyBase{tag, classification, section, label, help, ""},
    components[0], &properties.PropertyOptionsComponent{[]string{tag}, components, false},
  }
}

func (this *LumavateController) GetDynamicComponentsProperty(tag, name, classification, section, label, help string) *properties.PropertyComponents {

  components := this.GetComponentsWithTag(tag)

  if len(components) == 0 {
    return &properties.PropertyComponents{
      &properties.PropertyBase{name, classification, section, label, help, ""},
      []*properties.Component{}, &properties.PropertyOptionsComponents{[]string{}, []*properties.Component{}},
    }
  }

  return &properties.PropertyComponents{
    &properties.PropertyBase{name, classification, section, label, help, ""},
    []*properties.Component{}, &properties.PropertyOptionsComponents{[]string{tag}, components},
  }
}

func (this *LumavateController) InitBranding(instanceData *models.AppSettingsStruct){
  if instanceData.BodyMaxWidth != 0 {
    instanceData.BodyMaxWidthStr = fmt.Sprintf("%vpx", instanceData.BodyMaxWidth)
  } else {
    instanceData.BodyMaxWidthStr = "100%"
  }
}

func (this *LumavateController) InitFontStyles(instanceData *models.AppSettingsStruct) [] models.FontStyleDisplayStruct {
  styles := [] models.FontStyleDisplayStruct{}
	if instanceData.H1FontStyle != nil {
  	styles = append(styles, this.initFontStyle("h1", instanceData.H1FontStyle))
	}

	if instanceData.H2FontStyle != nil {
  	styles = append(styles, this.initFontStyle("h2", instanceData.H2FontStyle))
	}

	if instanceData.H3FontStyle != nil {
 		styles = append(styles, this.initFontStyle("h3", instanceData.H3FontStyle))
	}
	if instanceData.H4FontStyle != nil {
  	styles = append(styles, this.initFontStyle("h4", instanceData.H4FontStyle))
	}
	if instanceData.ParagraphFontStyle != nil {
 		styles = append(styles, this.initFontStyle("paragraph", instanceData.ParagraphFontStyle))
	}
	if instanceData.LinkFontStyle != nil {
 		styles = append(styles, this.initFontStyle("link", instanceData.LinkFontStyle))
	}
	if instanceData.ButtonFontStyle != nil {
 		styles = append(styles, this.initFontStyle("button", instanceData.ButtonFontStyle))
	}
  return styles

}

// Lets find a way to do this better.  Maybe we go back to do base components tied to page layout and flip based on mode
// For now, it's a hardcoded list of Lumavate Component Sets main DirectIncludes file names.•
func (this *LumavateController) ContainsIonicLibrary(includes []string) bool{

  for _, path := range includes {
    path_split := strings.Split(path, "/")
    file := path_split[len(path_split)-1]
    if strings.Contains(file, "luma-ion"){
      return true
    }
  }
  return false

}

func (this *LumavateController) initFontStyle(key string, fontStyle *models.FontStyleStruct) models.FontStyleDisplayStruct{
  underlineValue := "none"

  if fontStyle.FontUnderline {
    underlineValue = "underline"
  }

  return models.FontStyleDisplayStruct{
    FontColor: this.trimVar(fontStyle.FontColor),
    FontFamily: this.trimVar(fontStyle.FontFamily),
    FontSize: fmt.Sprint(fontStyle.FontSize, "px"),
    FontUnderline: underlineValue,
    Name: key,
  }
}

// Removes the var(--x) wrapper.
// TODO: Convert to regex if time allows
func (this *LumavateController) trimVar(cssVarRef string) string {
  trimmedValue := strings.TrimPrefix(cssVarRef, "var(--")
  trimmedValue = strings.TrimSuffix(trimmedValue, ")")
  return trimmedValue
}
