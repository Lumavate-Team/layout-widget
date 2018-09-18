package models

import (
  "fmt"
  common "github.com/Lumavate-Team/lumavate-go-common"
  widget "github.com/Lumavate-Team/lumavate-go-common/models"
  component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type MainController struct {
  common.LumavateController
}

type Footer struct {
  ComponentHtml string
  ComponentType string
}

type Modal struct {
  ComponentHtml string
  ComponentType string
}

type Header struct {
  ComponentHtml string
  ComponentType string
}

type BodyOptions struct {
  ComponentHtml string
  ComponentType string
  ComponentData struct {
    BodyTemplateColumns string
    BodyTemplateRows    string
    BodyRowGap          string
    BodyColumnGap       string
    JustifyContent      string
    AlignContent        string
    BodyNumRows         int
    BodyNumColumns      int
    BodyMaxWidth        int
    BodyMaxWidthStr     string
  }
}

type LumavateDomain struct {
  Payload struct {
    Data struct {
      Domain string
      RuntimeData map[string]interface{}
    }
  }
}

type LumavateRequest struct {
  Payload struct {
    Data struct {
      widget.CommonWidgetStruct
      DisplayBackgroundImage bool
      BackgroundImage        component_data.ImageStruct
      BackgroundColor        string
      DisplayHeader          bool
      DisplayFooter          bool
      BodyProperties         BodyOptions
      BodyItems              []LayoutContainer
      Footer                 widget.Component
      Header                 widget.Component
      ModalItems             []widget.Component
			LogicItems						 []LogicContainer
    }
  }
}

type LogicContainer struct {
	ComponentData struct {
		Placement string	
	}
	ComponentHtml string
}


type LayoutContainer struct {
  ComponentData struct {
    TemplateRowStart    int
    TemplateRowSpan     int
    TemplateColumnStart int
    TemplateColumnSpan  int
    CssClass            string
    DisplayMode         string
    JustifySelf         string
    AlignSelf           string
  }
  ComponentHtml string
}

func (this LogicContainer) GetHtml() string {
	return fmt.Sprintf(`
	%v
	`,
		this.ComponentHtml)
}

func (this LayoutContainer) GetHtml() string {
  return fmt.Sprintf(`
  <div class="layout-%v %v"
  style="justify-self:%v;align-self:%v;grid-area:%v/%v/ span %v/ span %v">
  %v
  </div>`,
    this.ComponentData.DisplayMode,
    this.ComponentData.CssClass,
    this.ComponentData.JustifySelf,
    this.ComponentData.AlignSelf,
    this.ComponentData.TemplateRowStart,
    this.ComponentData.TemplateColumnStart,
    this.ComponentData.TemplateRowSpan,
    this.ComponentData.TemplateColumnSpan,
    this.ComponentHtml)
}
