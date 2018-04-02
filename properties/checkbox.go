package properties

import (
  "encoding/json"
)

type PropertyCheckbox struct {
	*PropertyBase
  Default bool `json:"default"`
}

func (p *PropertyCheckbox) MarshalJSON() (b []byte, e error) {
  type Copy PropertyCheckbox

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"checkbox",
		(*Copy)(p),
	})
}
