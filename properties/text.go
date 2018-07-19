package properties

import (
  "encoding/json"
)

type PropertyOptionsText struct {
	ReadOnly bool `json:"readOnly"`
	Rows int `json:"rows"`
}

type PropertyText struct {
	*PropertyBase
	Default string `json:"default"`
	Options PropertyOptionsText `json:"options"`
}

func (p *PropertyText) MarshalJSON() (b []byte, e error) {
  type Copy PropertyText

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"text",
		(*Copy)(p),
	})
}

type PropertyTranslatedText struct {
	*PropertyBase
	Default string `json:"default"`
  Options PropertyOptionsText `json:"options"`
}

func (p *PropertyTranslatedText) MarshalJSON() (b []byte, e error) {
  type Copy PropertyTranslatedText

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"translated-text",
		(*Copy)(p),
	})
}

