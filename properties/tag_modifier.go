
package properties

import (
  "encoding/json"
)

type TagModifierOptions struct {
  TagType string `json:"tagType"`
  Properties []*PropertyType `json:"properties"`
}

type PropertyTagModifier struct {
  *PropertyBase
  Options TagModifierOptions `json:"options"`
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

