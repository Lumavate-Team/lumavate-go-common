package properties

import (
  "encoding/json"
)

type PropertyCodeEditor struct {
	*PropertyBase
	Default string `json:"default"`
}

func (p *PropertyCodeEditor) MarshalJSON() (b []byte, e error) {
  type Copy PropertyColor

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"code-editor",
		(*Copy)(p),
	})
}

