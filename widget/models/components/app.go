package components

import (
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"fmt"
	"encoding/json"
)

type AppStruct struct {
	ComponentData struct {
		Title string
		OpenNewWindow bool
		AppLink string
		Apple string
		Google string
		Microsoft string
		UseBackgroundColor bool
		BackgroundColor string
		UseBackgroundImage bool
		Image component_data.ImageStruct
		ImageScaling string
	}
}

func (this AppStruct) GetHtml() string {
	var style = ""
	if this.ComponentData.UseBackgroundColor != false {
		style = fmt.Sprintf(`background-color:%v;`,this.ComponentData.BackgroundColor)
	}

	if this.ComponentData.Image.Preview != "" {
		if this.ComponentData.UseBackgroundImage != false {
			style = style + fmt.Sprintf(`background-image:url('%v')`,this.ComponentData.Image.Preview)
		}
	}

	data, _ := json.Marshal(this.ComponentData)
	return fmt.Sprintf(`
		<div class="layout-nav-item layout-nav-tile layout-%v" style="%v" onclick='appNavigate(%s,event)' > </div>`,
		this.ComponentData.ImageScaling,
		style,
		data)
}
