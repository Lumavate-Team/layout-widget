package components

import (
	"fmt"
)

type TextStruct struct {
	ComponentData struct {
		Title string
		Text string
	}
}

func (this TextStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div class="nav-item nav-tile">
				%v
		</div>`,
		this.ComponentData.Text)
}
