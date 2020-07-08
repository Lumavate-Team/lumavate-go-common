package common

import (
	models "github.com/Lumavate-Team/lumavate-go-common/models"
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"html/template"
  "strings"
  "regexp"
)

func SafeCss(in string) (out template.CSS) {
	out = template.CSS(in)
	return
}

func SafeHtml(in string) (out template.HTML) {
	out = template.HTML(in)
	return
}

func HasSuffix(in string, test string) (out bool){
  return strings.HasSuffix(in, test)
}
func HasPrefix(in string, test string) (out bool){
  return strings.HasPrefix(in, test)
}
func Replace(input, from,to string) string {
 \return strings.Replace(input,from,to, -1)
}

func Esm(src string)(out string){
  extIndex := strings.LastIndex(src, ".js")
  return src[:extIndex] + ".esm.js"
}


func EscapeSpecial(src string)string{
  reg, _:= regexp.Compile("[^a-zA-Z0-9_-]+")
  s := strings.Replace(src, " ", "_", -1)
  return reg.ReplaceAllString(s, "")
}

func ComponentHtml(in component_data.ComponentData) (out template.HTML) {
	out = template.HTML(in.GetHtml())
	return
}

func ModalHtml(in models.ComponentStruct) (out template.HTML) {
	out = template.HTML(in.GetHtml())
	return
}
