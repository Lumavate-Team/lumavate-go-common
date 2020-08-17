package models

type AppSettingsStruct struct {
  InstanceName      string                        `json:"instance__name"`
  PageType          PageTypeStruct  `json:"pageType"`
  DirectIncludes    [] string                     `json:"__directIncludes"`
  DirectCssIncludes [] string                     `json:"__directCssIncludes"`
  Header ComponentStruct `json:"header"`
  Footer ComponentStruct `json:"footer"`
  Fab ComponentStruct `json:"fab"`

  BodyMaxWidth int `json:"bodyMaxWidth"`
  BodyMaxWidthStr string `json:"bodyMaxWidthStr"`
  ThemeDataStruct
}
