package models

type SecurityDataStruct struct {
  ComponentStruct
  ComponentData struct {
    NoAuthRedirect PageLinkStruct
    SpecificGroup [] string `json:"specificGroup"`
  }
}
