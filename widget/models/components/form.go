package components

import (
	"fmt"
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type FormStruct struct {
	ComponentData struct {
		Title string
		Text string
		FormInputs component_data.FormTextStruct
	}
}

func (this FormStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div class="nav-item nav-tile">
				%v
				%v
				%v
		</div>`,
		this.ComponentData.Text,
		this.ComponentData.FormInputs.GetHtml(),
		this.ComponentData.FormInputs)
}