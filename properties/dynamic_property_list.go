package properties

import (
  "encoding/json"
)

type DynamicPropertyOptions struct {
	PropertyDef []byte `json:"propertyDef"`
	OnSettings bool `json:"onSettings"`
	Max int `json:"max"`
	FullHeight bool `json:"fullHeight"`
}

type DynamicPropertyList struct {
  *PropertyBase
  Options DynamicPropertyOptions `json:"options"`
}

func (p *DynamicPropertyList) MarshalJSON() (b []byte, e error) {
  type Copy DynamicPropertyList

  return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"dynamic-property-list",
		(*Copy)(p),
	})
}