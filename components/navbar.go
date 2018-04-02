package components

import (
  "properties"
  "properties/component_data"
  "fmt"
  "encoding/json"
)

/*
 * Gets a property that allows entry for "NavBar" data
 */
func GetNavBarProperty() *properties.PropertyComponent {
  return &properties.PropertyComponent{
    &properties.PropertyBase{"navBar", "Nav Bar", "Nav Bar Properties", "Nav Bar", ""},
    GetNavBarComponent(), properties.PropertyOptionsComponent{[] string {"navBar"}, [] *properties.Component {GetNavBarComponent()} },
  }
}

/*
 * Gets a property that allows entry for 'NavBarItems' data
 */
func GetNavBarItemsProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents{
    &properties.PropertyBase{"navBarItems", "Nav Bar", "Nav Bar Items", "Nav Bar", ""},
    [] *properties.Component {}, properties.PropertyOptionsComponent{[] string {"navBarItem"}, [] *properties.Component {GetNavBarItemComponent()} },
  }
}

/*
 * Gets a description for the 'NavBar' component.  This is defined in a central place
 */
func GetNavBarComponent() *properties.Component {
  //return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
        comp := properties.LoadComponent("https://experience.john.labelnexusdev.com", "1.0.0", "navBar")
        comp.Category = "navBar"
        return comp
}

/*
 * Gets a description for the 'NavBarItem' component.  This is defined in a central place
 */
func GetNavBarItemComponent() *properties.Component {
  //return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
        comp :=properties.LoadComponent("https://experience.john.labelnexusdev.com", "1.0.0", "navBarItem")
        comp.Category = "navBarItem"
        return comp
}

type NavBarItemsStruct struct {
  ComponentData struct {
    Items []NavBarItemStruct `json:"navBarItem"`
  }
}

type NavBarItemStruct struct {
  ComponentData struct {
    Text string
    LinkTo component_data.PageLinkStruct
    ImageSource component_data.ImageStruct
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
    LinkTo component_data.PageLinkStruct `json:"linkTo"`
    ImageLink component_data.ImageStruct `json:"imageSource"`
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
