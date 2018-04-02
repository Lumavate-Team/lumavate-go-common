package properties

import (
  "encoding/json"
)

type PropertyImage struct {
	*PropertyBase
}

func (p *PropertyImage) MarshalJSON() (b []byte, e error) {
  type Copy PropertyImage

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"image-upload",
		(*Copy)(p),
	})
}

