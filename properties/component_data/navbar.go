package component_data

import (
  "fmt"
  "encoding/json"
)

type NavBarItemStruct struct {
  ComponentData struct {
    Text string
    LinkTo PageLinkStruct
    ImageSource ImageStruct
  }
}

type NavBarItemsStruct struct {
  ComponentData struct {
    Items []NavBarItemStruct `json:"navBarItem"`
  }
}

type NavBarStruct struct {
  ComponentData struct {
    ShowNavBar bool
    ItemColor string
    BackgroundColor string
    Position string
    NavBarItems NavBarItemsStruct `json: "-"`
  }
}

type NavBarContainerStruct struct {
  NavBar NavBarStruct `json: "navBar"`
  NavBarItems NavBarItemsStruct `json: "navBarItems"`
}

func (m *NavBarContainerStruct) UnmarshalJSON(bytes []byte) error{
  var tmp NavBarContainerStruct

  if err := json.Unmarshal(bytes, &tmp); err != nil {
    return nil, err
  }
  m.NavBar = tmp.NavBar
  m.NavBar.ComponentData.NavBarItems = tmp.NavBarItems
  m.NavBarItems = tmp.NavBarItems
  return nil
}
func (m *NavBarItemsStruct) UnmarshalJSON(bytes []byte) error {
  var tmp []NavBarItemStruct

  if err := json.Unmarshal(bytes, &tmp); err != nil {
    return err
  }

  m.ComponentData.Items=tmp
  return nil
}

func (m *NavBarItemStruct) MarshalJSON() ([]byte, error) {
  return json.Marshal(&struct {
    Text string `json:"text"`
    LinkTo PageLinkStruct `json:"linkTo"`
    ImageLink ImageStruct `json:"imageSource"`
  }{
    Text: m.ComponentData.Text,
    LinkTo: m.ComponentData.LinkTo,
    ImageLink: m.ComponentData.ImageSource,
  })
}

func (this NavBarStruct) GetHtml() string {
  nav := `<lumavate-nav-bar
      nav-bar-background-color='%v'
      nav-bar-item-color='%v'
      nav-bar-position='%v'
      nav-bar-show-nav-bar='%v'
      nav-bar-items='%v'>
    </lumavate-nav-bar>`

    items, err := json.Marshal(this.ComponentData.NavBarItems.ComponentData.Items)
    if err != nil {
      fmt.Println(err)
    }

    return fmt.Sprintf(nav, 
      this.ComponentData.BackgroundColor, 
      this.ComponentData.ItemColor, 
      this.ComponentData.Position,
      this.ComponentData.ShowNavBar,
      string(items))
}
