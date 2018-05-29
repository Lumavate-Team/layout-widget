package components

import (
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"fmt"
)

type NavigationStruct struct {
	ComponentData struct {
		Title string
		UseBackgroundColor bool
		BackgroundColor string
		UseBackgroundImage bool
		Image component_data.ImageStruct
		ImageScaling string
		PageLink component_data.PageLinkStruct
	}
}

func (this NavigationStruct) GetHtml() string {
	var style = ""
	if this.ComponentData.UseBackgroundColor != false {
		style = fmt.Sprintf(`background-color:%v;`,this.ComponentData.BackgroundColor)
	}

	if this.ComponentData.Image.Preview != "" {
		if this.ComponentData.UseBackgroundImage != false {
			style = style + fmt.Sprintf(`background-image:url('%v')`,this.ComponentData.Image.Preview)
		}

		if this.ComponentData.PageLink.Url != "" {
			return fmt.Sprintf(`
			<div class="layout-nav-item layout-nav-tile layout-%v" onclick="navigate('%v')" style="%v"> </div>`,
				this.ComponentData.ImageScaling,
				this.ComponentData.PageLink.Url,
				style)
		} else {
			return fmt.Sprintf(`
			<div class="layout-nav-item layout-nav-tile layout-%v" style="%v"> </div>`,
				this.ComponentData.ImageScaling,
				style)
		}
	} else {
		if this.ComponentData.PageLink.Url != "" {
			return fmt.Sprintf(`
			<div class="layout-nav-item" onclick="navigate('%v')" style="%v">
						%v
				</div>`,
				this.ComponentData.Title,
				style,
				this.ComponentData.Title)
		} else {
			return fmt.Sprintf(`
			<div class="layout-nav-item" onclick="navigate('%v')" style="%v">
						%v
				</div>`,
				this.ComponentData.Title,
				style,
				this.ComponentData.Title)
		}
	}
}
