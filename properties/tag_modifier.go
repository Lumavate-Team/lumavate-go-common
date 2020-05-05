
package properties

import (
  "encoding/json"
)

type PropertyTagModifier struct {
  TagName string `json:"tagName"`
	Properties [] PropertyType `json:"properties"`
}

func (p *PropertyTagModifier) MarshalJSON() (b []byte, e error) {
  type Copy PropertyTagModifier

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"tag-modifier",
		(*Copy)(p),
	})
}

