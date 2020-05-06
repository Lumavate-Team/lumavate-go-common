package properties

import (
  "encoding/json"
)

type DynamicOptions struct {
  ComponentTagName string `json:"componentTagName"`
  AllowMultiple bool `json:"allowMultiple"`
	Components [] *Component `json:"components"`
}
type DynamicComponent struct {
  *PropertyBase
	Options *DynamicOptions `json:"options"`
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

