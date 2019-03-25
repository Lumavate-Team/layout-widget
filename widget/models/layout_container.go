package models

import (
	"fmt"
  common_models "github.com/Lumavate-Team/lumavate-go-common/models"
)

type LayoutContainerStruct struct {
	common_models.ComponentStruct
	ComponentData struct {
		TemplateRowStart    int
		TemplateRowSpan     int
		TemplateColumnStart int
		TemplateColumnSpan  int
		DisplayMode         string
		JustifySelf         string
		AlignSelf           string
	}
}

func (this LayoutContainerStruct) GetHtml() string {
  return fmt.Sprintf(`
  <div class="layout-%v"
	style="justify-self:center;align-self:start;grid-area:%v/%v/ span %v/ span %v;width:%v;height:%v">
  %v
  </div>`,
    this.ComponentData.DisplayMode,
    this.ComponentData.TemplateRowStart,
    this.ComponentData.TemplateColumnStart,
    this.ComponentData.TemplateRowSpan,
    this.ComponentData.TemplateColumnSpan,
		"100%",
		"100%",
    this.ComponentHtml)
}
