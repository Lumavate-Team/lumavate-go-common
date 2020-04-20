
package models

type AppSettingsStruct struct {
  PageType PageTypeStruct `json:"pageType"`
  Footer ComponentStruct `json:"footer"`
  Header ComponentStruct `json:"header"`
  Fab ComponentStruct `json:"fab"`
  //AddToHome AddToHomeStruct `json:"addToHome"`

  // TODO: Move Add to Home properties into actual PropertyComponent instead of flat list of properties
  // and add custom unmarshal to move these properties if possible(would break page builder though)
  BodyMaxWidth int `json:"bodyMaxWidth"`
  BodyMaxWidthStr string `json:"bodyMaxWidthStr"`
  ShowAddToHome bool `json:"showAddToHome"`
  SkipFirst bool `json:"skipFirst"`
  StartDelay int `json:"startDelay"`
  Lifespan int `json:"lifespan"`
  DisplayCount int `json:"displayCount"`
  Message string `json:"message"`

  StyleData [] Styles `json:"styleData"`
}
