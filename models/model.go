package models

import (
  "properties/component_data"
  "components"
)

type WidgetStruct struct {
    PageType component_data.PageTypeStruct
    BackgroundColor string
    components.NavBarContainerStruct
}
