package models

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type CommonWidgetStruct struct {
  InstanceName string `json:"instance__name"`
  PageType component_data.PageTypeStruct `json:"pageType"`
  component_data.NavBarContainerStruct
}


