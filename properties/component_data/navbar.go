package component_data

import (
  _ "fmt"
)

type NavBarStruct struct {
  ComponentData struct {
    ShowNavBar bool
    ItemColor string
    BackgroundColor string
    position string
  }
}

func (this NavBarStruct) GetHtml() string {
  return "Nav Bar Goes Here"
}

type NavBarItemStruct struct {
  ComponentData struct {
    Text string
    LinkTo PageLinkStruct
    ImageSource ImageStruct
  }
}

func (this NavBarItemStruct) GetHtml() string {
  return "Nav Bar Item Goes Here"
}

