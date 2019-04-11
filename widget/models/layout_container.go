package models

import (
	"fmt"
	common_models "github.com/Lumavate-Team/lumavate-go-common/models"
	//"html/template"
)

type LayoutContainerStruct struct {
	common_models.ComponentStruct
	ComponentData struct {
		TemplateRowStart    int
		TemplateRowSpan     int
		TemplateColumnStart int
		TemplateColumnSpan  int
		DisplayMode         string
		HorizontalAlignment string
		VerticalAlignment   string
		PaddingTop          int
		PaddingRight        int
		PaddingLeft         int
		PaddingBottom       int
	}
}

func (this LayoutContainerStruct) GetHtml() string {
	return fmt.Sprintf(`
  <div class="layout-%v"
	style="justify-self:%v;align-self:%v;grid-area:%v/%v/ span %v/ span %v;padding-top:%vpx;padding-right:%vpx;padding-bottom:%vpx;padding-left:%vpx">
  %v
  </div>`,
		this.ComponentData.DisplayMode,
		this.ComponentData.HorizontalAlignment,
		this.ComponentData.VerticalAlignment,
		this.ComponentData.TemplateRowStart,
		this.ComponentData.TemplateColumnStart,
		this.ComponentData.TemplateRowSpan,
		this.ComponentData.TemplateColumnSpan,
		this.ComponentData.PaddingTop,
		this.ComponentData.PaddingRight,
		this.ComponentData.PaddingBottom,
		this.ComponentData.PaddingLeft,
		this.ComponentHtml)
	//template.HTML("<div style='display:inline-block;text-align:center;width:100%;height:100%;border:1px dotted red'>HELLO</div>"))
}
