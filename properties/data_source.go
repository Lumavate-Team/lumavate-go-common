package properties

import (
  "encoding/json"
)

type DataSource struct {
  *PropertyBase
}

func (p *DataSource) MarshalJSON() (b []byte, e error) {
  type Copy DataSource

  return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"data-source",
		(*Copy)(p),
	})
}