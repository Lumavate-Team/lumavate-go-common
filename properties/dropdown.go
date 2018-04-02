package properties

import (
  "encoding/json"
)

type PropertyDropdown struct {
	*PropertyBase
	Default string `json:"default"`
	Options map[string]string `json:"options"`
}

func (p *PropertyDropdown) MarshalJSON() (b []byte, e error) {
  type Copy PropertyDropdown

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"dropdown",
		(*Copy)(p),
	})
}
