package properties

import (
  "encoding/json"
)

type PropertyToggle struct {
	*PropertyBase
	Default bool `json:"default"`
}

func (p *PropertyToggle) MarshalJSON() (b []byte, e error) {
  type Copy PropertyToggle

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"toggle",
		(*Copy)(p),
	})
}
