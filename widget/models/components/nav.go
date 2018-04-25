package components

import (
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"fmt"
)

type NavigationStruct struct {
	ComponentData struct {
		Title string
		Image component_data.ImageStruct
		ImageScaling string
		PageLink component_data.PageLinkStruct
	}
}

func (this NavigationStruct) GetHtml() string {
	if this.ComponentData.Image.Preview != "" {
		if this.ComponentData.PageLink.Url != "" {
			return fmt.Sprintf(`
				<div class="nav-item nav-tile %v" onclick="navigate('%v')" style="background-image:url('%v');"> </div>`,
				this.ComponentData.ImageScaling,
				this.ComponentData.PageLink.Url,
				this.ComponentData.Image.Preview)
		} else {
			return fmt.Sprintf(`
				<div class="nav-item nav-tile %v" style="background-image:url('%v');"> </div>`,
				this.ComponentData.ImageScaling,
				this.ComponentData.Image.Preview)
		}
	} else {
		if this.ComponentData.PageLink.Url != "" {
			return fmt.Sprintf(`
				<div class="nav-item" onclick="navigate('%v')">
						%v
				</div>`,
				this.ComponentData.Title,
				this.ComponentData.Title)
		} else {
			return fmt.Sprintf(`
				<div class="nav-item" onclick="navigate('%v')">
						%v
				</div>`,
				this.ComponentData.Title,
				this.ComponentData.Title)
		}
	}
}