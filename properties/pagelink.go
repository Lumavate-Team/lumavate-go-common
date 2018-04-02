package properties

import (
  "encoding/json"
)

type PropertyPageLink struct {
  *PropertyBase
}

func (p *PropertyPageLink) MarshalJSON() (b []byte, e error) {
  type Copy PropertyPageLink

  return json.Marshal(&struct{
    Type string `json:"type"`
    *Copy
  }{
    "page-link",
    (*Copy)(p),
  })
}
