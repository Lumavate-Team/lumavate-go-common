
package models

type AppSettingsStruct struct {
  Footer ComponentStruct `json:"footer"`
  Header ComponentStruct `json:"header"`
  Fab ComponentStruct `json:"fab"`
  //AddToHome AddToHomeStruct `json:"addToHome"`

  // TODO: Move Add to Home properties into actual PropertyComponent instead of flat list of properties
  // and add custom unmarshal to move these properties if possible(would break page builder though)
  BodyMaxWidth int `json:"bodyMaxWidth"`
  BodyMaxWidthStr string `json:"bodyMaxWidthStr"`
  StyleData [] Styles `json:"styleData"`
  AddToHomeStruct
}
