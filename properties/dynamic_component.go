package properties

import (
  "encoding/json"
)

type DynamicOptions struct {
  Categories [] string `json:"categories"`
	Components [] *Component `json:"components"`
  TagModifiers []PropertyType `json:"tagModifiers"`
  Position string `json:"position"`
  IsPrimary bool `json:"isPrimary"`
  HeaderColor string `json:"headerColor"`
}
type DynamicComponentOptions struct {
  DynamicOptions
  Required bool `json:"isRequired"`
}

type DynamicComponent struct {
  *PropertyBase
	Options *DynamicComponentOptions `json:"options"`
}

func (p *DynamicComponent) MarshalJSON() (b []byte, e error) {
  type Copy DynamicComponent

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"dynamic-component",
		(*Copy)(p),
	})
}

type DynamicComponents struct {
  *PropertyBase
  Options *DynamicOptions `json:"options"`
}

func (p *DynamicComponents) MarshalJSON() (b []byte, e error) {
  type Copy DynamicComponents

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
    "dynamic-components",
		(*Copy)(p),
	})
}
