package models

import (
	"fmt"
)

type ComponentStruct struct {
	ComponentHtml string
	ComponentType string
}

func (this ComponentStruct) GetHtml() string {
	return fmt.Sprintf(`<div>%v</div>=`, this.ComponentHtml)
}
