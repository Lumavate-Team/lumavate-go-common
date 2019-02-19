package properties

import (
//	_ "github.com/Lumavate-Team/lumavate-go-common"
)

type LumavateProperties struct {
	Authorization     string
	DynamicComponents DynamicComponents
	BaseUrl           string
}

func NewLumavateProperties(auth string, base_url string) *LumavateProperties {
	lp := &LumavateProperties{auth, DynamicComponents{}, base_url}
	lp.DynamicComponents.LoadAllComponentSets(lp.Authorization, base_url)
	return lp
}
