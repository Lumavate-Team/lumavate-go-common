package models

type AppSettingsStruct struct {
  InstanceName      string                        `json:"instance__name"`
  PageType          PageTypeStruct  `json:"pageType"`
  DirectIncludes    [] string                     `json:"__directIncludes"`
  DirectCssIncludes [] string                     `json:"__directCssIncludes"`
  Header ComponentStruct `json:"header"`
  Footer ComponentStruct `json:"footer"`
  PoweredBy bool `json:"poweredBy"`
  Fab ComponentStruct `json:"fab"`
	CookieManagement ComponentStruct `json:"cookieManagement"`

  BodyMaxWidth int `json:"bodyMaxWidth"`
  BodyMaxWidthStr string `json:"bodyMaxWidthStr"`
  ThemeDataStruct
}
