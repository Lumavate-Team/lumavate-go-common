package properties

import (
  "encoding/json"
)
type CodeEditorOptions struct {
	Language string `json:"language"`
}

type PropertyCodeEditor struct {
	*PropertyBase
	Default string `json:"default"`
	Options CodeEditorOptions `json:"options"`
}

func (p *PropertyCodeEditor) MarshalJSON() (b []byte, e error) {
  type Copy PropertyCodeEditor

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"code-editor",
		(*Copy)(p),
	})
}

