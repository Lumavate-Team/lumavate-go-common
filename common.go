package common

import (
	models "github.com/Lumavate-Team/lumavate-go-common/models"
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"html/template"
)

func SafeCss(in string) (out template.CSS) {
	out = template.CSS(in)
	return
}

func SafeHtml(in string) (out template.HTML) {
	out = template.HTML(in)
	return
}

func ComponentHtml(in component_data.ComponentData) (out template.HTML) {
	out = template.HTML(in.GetHtml())
	return
}

func ModalHtml(in models.Component) (out template.HTML) {
	out = template.HTML(in.GetHtml())
	return
}
