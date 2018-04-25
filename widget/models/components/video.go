package components

import (
	"fmt"
)

type VideoStruct struct {
	ComponentData struct {
		Title string
		Video  string
	}
}

func (this VideoStruct) GetHtml() string {
	return fmt.Sprintf(`
		<div class="nav-item nav-tile">
			<iframe class="video-frame" src="%v" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen ></iframe>
		</div>`,
		this.ComponentData.Video)
}