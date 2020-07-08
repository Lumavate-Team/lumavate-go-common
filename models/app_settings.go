package models

type AppSettingsStruct struct {
  ThemeDataStruct
  Header ComponentStruct `json:"header"`
  Footer ComponentStruct `json:"footer"`
  Fab ComponentStruct `json:"fab"`


  BodyMaxWidth int `json:"bodyMaxWidth"`
  BodyMaxWidthStr string `json:"bodyMaxWidthStr"`
}
