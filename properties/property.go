package properties

type PropertyType interface {
}

type PropertyBase struct {
	Name string `json:"name"`
	Classification string `json:"classification"`
	Section string `json:"section"`
	Label string `json:"label"`
	HelpText string `json:"helpText"`
  	HelpId string `json:"helpId,omitempty"`
}

type TypePropertyBase struct {
	*PropertyBase
	Type string `json:"type"`
}


