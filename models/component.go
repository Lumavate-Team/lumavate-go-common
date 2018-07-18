package models

import (
	"fmt"
)

type Component struct {
	ComponentHtml string
	ComponentType string
}

func (this Component) GetHtml() string {
	return fmt.Sprintf(`
  <div>%v</div>
  `,
		this.ComponentHtml)
}
