package properties

import (
  "encoding/json"
  "net/http"
  "io/ioutil"
	"fmt"
)

func LoadComponent(base_uri string, version string, component string) *Component {
	uri := fmt.Sprintf("%s/iot/v1/common-components/%s/%s", base_uri, version, component)
  req, _ := http.NewRequest("GET", uri, nil)
  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  c := Component{}
  json.Unmarshal(body, &c)
	return &c
}

type ComponentOptions struct {
	Categories [] string `json"categories"`
	Components [] Component `json"components"`
}

type Component struct {
	Category string `json:"category"`
	Section string `json:"section"`
	Type string `json:"type"`
	DisplayName string `json:"displayName"`
	Icon string `json:"icon"`
	Label string `json:"label"`
	Properties [] PropertyType `json:"properties"`
	DisplayNameTemplate string `json:"displayNameTemplate"`
}

type PropertyOptionsComponent struct {
	Categories [] string `json:"categories"`
	Components [] *Component `json:"components"`
}

type PropertyComponent struct {
	*PropertyBase
	DefaultComponent *Component `json:"-"`
	Options *PropertyOptionsComponent `json:"options"`
}

func (p *PropertyComponent) MarshalJSON() (b []byte, e error) {
  type Copy PropertyComponent

	type DefaultStruct struct {
		ComponentType string `json:"componentType"`
	}

  type MarshalledPropertyComponent struct {
		Type string `json:"type"`
    *Copy
    Default DefaultStruct `json:"default"`
	}

	return json.Marshal(&MarshalledPropertyComponent{
		"component",
		(*Copy)(p),
		DefaultStruct{ p.DefaultComponent.Type },
	})
}

type PropertyComponents struct {
	*PropertyBase
	Default [] *Component `json:"default"`
	Options *PropertyOptionsComponent `json:"options"`
}

func (p *PropertyComponents) MarshalJSON() (b []byte, e error) {
  type Copy PropertyComponents

	return json.Marshal(&struct{
		Type string `json:"type"`
		*Copy
	}{
		"components",
		(*Copy)(p),
	})
}

