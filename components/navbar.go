package components

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties"
)

/*
 * Gets a property that allows entry for "NavBar" data
 */
func GetNavBarProperty() *properties.PropertyComponent {
  return &properties.PropertyComponent{
    &properties.PropertyBase{"navBar", "Nav Bar", "", "Nav Bar", "", ""},
    GetNavBarComponent(), &properties.PropertyOptionsComponent{[] string {"navBar"}, [] *properties.Component {GetNavBarComponent()}, false },
  }
}

/*
 * Gets a property that allows entry for 'NavBarItems' data
 */
func GetNavBarItemsProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents{
    &properties.PropertyBase{"navBarItems", "Nav Bar", "Nav Bar Items", "Nav Bar", "", ""},
    [] *properties.Component {}, &properties.PropertyOptionsComponent{[] string {"navBarItem"}, [] *properties.Component {GetNavBarItemComponent()}, false },
  }
}

/*
 * Gets a description for the 'NavBar' component.  This is defined in a central place
 */
func GetNavBarComponent() *properties.Component {
  //return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
        comp := properties.LoadComponent("https://experience.lumavate.com", "1.0.0", "navBar")
        comp.Category = "navBar"
        return comp
}

/*
 * Gets a description for the 'NavBarItem' component.  This is defined in a central place
 */
func GetNavBarItemComponent() *properties.Component {
  //return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
        comp :=properties.LoadComponent("https://experience.lumavate.com", "1.0.0", "navBarItem")
        comp.Category = "navBarItem"
        return comp
}
