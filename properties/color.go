package properties

import (
  "encoding/json"
)

type PropertyColor struct {
	*PropertyBase
	Default string `json:"default"`
}

func (p *PropertyColor) MarshalJSON() (b []byte, e error) {
  type Copy PropertyColor

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"color",
		(*Copy)(p),
	})
}

