package properties

import (
  "encoding/json"
)

type PropertyOptionsNumeric struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type PropertyNumeric struct {
	*PropertyBase
	Default int `json:"default"`
  Options PropertyOptionsNumeric `json:"options"`
}

func (p *PropertyNumeric) MarshalJSON() (b []byte, e error) {
  type Copy PropertyNumeric

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"numeric",
		(*Copy)(p),
	})
}
