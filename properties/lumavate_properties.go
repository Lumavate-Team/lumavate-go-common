package properties

import (
//	_ "github.com/Lumavate-Team/lumavate-go-common"
)

type LumavateProperties struct {
	Authorization     string
	DynamicComponents DynamicComponents
}

func NewLumavateProperties(auth string) *LumavateProperties {
	lp := &LumavateProperties{auth, DynamicComponents{}}
	lp.DynamicComponents.LoadAllComponentSets(lp.Authorization, "")
	return lp
}
