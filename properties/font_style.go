
package properties

import (
  "encoding/json"
)

type FontStyleOptions struct {
	IncludeUnderline bool `json:"includeUnderline"`
}

type FontStyleDefault struct {
	FontFamily string `json:"fontFamily"`
	FontSize int `json:"fontSize"`
	FontColor string `json:"fontColor"`
}

type PropertyFontStyle struct {
	*PropertyBase
	Default FontStyleDefault `json:"default"`
	Options FontStyleOptions `json:"options"`
}

func (p *PropertyFontStyle) MarshalJSON() (b []byte, e error) {
  type Copy PropertyFontStyle

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"font-style",
		(*Copy)(p),
	})
}

