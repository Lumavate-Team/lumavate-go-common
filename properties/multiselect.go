package properties

import (
  "encoding/json"
)

type MultiselectOption struct {
	Value string `json:"value"`
	DisplayValue string `json:"displayValue"`
}

type PropertyMultiselect struct {
	*PropertyBase
	Default [] string `json:"default"`
	Options [] MultiselectOption `json:"options"`
}

func (p *PropertyMultiselect) MarshalJSON() (b []byte, e error) {
  type Copy PropertyMultiselect

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"multiselect",
		(*Copy)(p),
	})
}
