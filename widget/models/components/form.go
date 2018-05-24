package components

import (
	"fmt"
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type FormStruct struct {
	ComponentData struct {
		Title string
		Text string
		FormInputs component_data.FormStruct
	}
}

func (lc *FormStruct) UnmarshalJSON(data []byte) error {

	lc.ComponentData.FormInputs.UnmarshalJSON(data)
	return nil
}

func (this FormStruct) GetHtml() string {

	var html string

	for _, element := range this.ComponentData.FormInputs.FormItems.ComponentData.Forms {
		html = html + fmt.Sprintf(`
			<div class="nav-item nav-tile">
				%v
				%v
				%v
			</div>`,
			this.ComponentData.Text,
			element.GetHtml(),
			this.ComponentData.FormInputs)
	}

	return fmt.Sprintf(`
		<form action="">
			%v
		</form>`,
		html)
}