package common

import (
  "encoding/json"
  "github.com/Lumavate-Team/lumavate-go-common/properties"
  "github.com/Lumavate-Team/lumavate-go-common/api_core"
)

type DynamicComponent struct {
  Icon string
  Label string
  Type string
  Tags [] string
  Template string
  Properties [] properties.PropertyType
}

type ComponentSetRequest struct {
  Payload struct {
    Data [] struct {
      CurrentVersion struct {
        DirectIncludes [] string
        DirectCssIncludes [] string
        Distribution string
        Components [] *DynamicComponent
      }
    }
  }
}

type DynamicComponents struct {
  Components [] *DynamicComponent
}

func (self *DynamicComponents) LoadAllComponentSets(authorization string) {
  lr := api_core.LumavateRequest{authorization}
  body, _ := lr.Get("/pwa/v1/component-sets")
  cs := ComponentSetRequest{}
  json.Unmarshal(body, &cs)

  for _, set := range cs.Payload.Data {
    for _, component := range set.CurrentVersion.Components {
      self.Components = append(self.Components, component)
    }
  }
}

func (self *DynamicComponents) GetComponentsWithTag(tag string) []*properties.Component {

  components := [] *properties.Component {}

  for _, component := range self.Components {
    for _, t := range component.Tags {
      if t == tag {
        components = append(components, &properties.Component{tag, "", component.Type, "", component.Icon, component.Label, component.Properties})
      }
    }
  }

  return components
}

func (self *DynamicComponents) GetDynamicComponentProperty(tag, name, classification, section, label, help string) *properties.PropertyComponent {

  components := self.GetComponentsWithTag(tag)

  if len(components) == 0 {
    return &properties.PropertyComponent{
      &properties.PropertyBase{tag, classification, section, label, help},
      &properties.Component{}, &properties.PropertyOptionsComponent{[] string {}, [] *properties.Component {} },
    }
  }

  return &properties.PropertyComponent{
    &properties.PropertyBase{tag, classification, section, label, help},
    components[0], &properties.PropertyOptionsComponent{[] string {tag}, components },
  }
}

func (self *DynamicComponents) GetDynamicComponentsProperty(tag, name, classification, section, label, help string) *properties.PropertyComponents {

  components := self.GetComponentsWithTag(tag)

  if len(components) == 0 {
    return &properties.PropertyComponents{
      &properties.PropertyBase{name, classification, section, label, help},
      [] *properties.Component{}, &properties.PropertyOptionsComponent{[] string {}, [] *properties.Component {} },
    }
  }

  return &properties.PropertyComponents{
    &properties.PropertyBase{name, classification, section, label, help},
    [] *properties.Component{}, &properties.PropertyOptionsComponent{[] string {tag}, components },
  }
}
