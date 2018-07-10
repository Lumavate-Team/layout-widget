package models

import (
  common "github.com/Lumavate-Team/lumavate-go-common"
  widget "github.com/Lumavate-Team/lumavate-go-common/models"
  component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
  "fmt"
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

type LumavateRequest struct {
  Payload struct {
    Data struct {
      widget.CommonWidgetStruct
      InlineCss string
      DisplayBackgroundImage bool
      BackgroundImage component_data.ImageStruct
      BackgroundColor string
      BodyTemplateColumns string
      BodyTemplateRows string
      BodyRowGap string
      BodyColumnGap string
      DisplayHeader bool
      DisplayFooter bool
      JustifyContent string
      AlignContent string
      BodyItems []LayoutContainer
      Footer Footer
      Header Header
      ModalItems []Modal
      DirectIncludes []string `json:"__directIncludes"`
    }
  }
}

func (this Modal) GetHtml() string {
  return fmt.Sprintf(`
  <div>%v</div>
  `,
  this.ComponentHtml)
}

type LayoutContainer struct {
  ComponentData struct {
    TemplateRowStart string
    TemplateRowEnd string
    TemplateColumnStart string
    TemplateColumnEnd string
    CssClass string
    DisplayMode string
    JustifySelf string
    AlignSelf string
  }
  ComponentHtml string
}

func (this LayoutContainer) GetHtml() string {
  return fmt.Sprintf(`
  <div class="layout-%v %v"
  style="justify-self:%v;align-self:%v;grid-area:%v/%v/%v/%v">
  %v
  </div>`,
  this.ComponentData.DisplayMode,
  this.ComponentData.CssClass,
  this.ComponentData.JustifySelf,
  this.ComponentData.AlignSelf,
  this.ComponentData.TemplateRowStart,
  this.ComponentData.TemplateColumnStart,
  this.ComponentData.TemplateRowEnd,
  this.ComponentData.TemplateColumnEnd,
  this.ComponentHtml)
}
