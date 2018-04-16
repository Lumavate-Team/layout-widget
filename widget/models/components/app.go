package components

import (
	"fmt"
)

type AppStruct struct {
	ComponentData struct {
		Title string
		App string
	}
}

func (this AppStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div class="nav-item nav-tile">
			%v
		</div>`,
		this.ComponentData.App)
}

