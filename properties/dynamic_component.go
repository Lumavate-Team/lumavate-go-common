

package properties

import (
  "encoding/json"
)

type DynamicComponent struct {
  *PropertyBase
  ComponentTagName string `json:"componentTagName"`
  AllowMultiple bool `json:"allowMultiple"`
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

