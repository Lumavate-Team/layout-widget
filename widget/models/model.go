package models

import (
	common "github.com/Lumavate-Team/lumavate-go-common"
	widget "github.com/Lumavate-Team/lumavate-go-common/models"
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	components "widget/models/components"
	"fmt"
	"encoding/json"
	"reflect"
)

type MainController struct {
  common.LumavateController
}

type Footer struct {
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
			GridTemplateColumns string
			GridTemplateRows string
      GridRowGap string
      GridColumnGap string
      JustifyContent string
      AlignContent string
			GridItems []LayoutContainer
      Footer Footer
      DirectIncludes []string `json:"__directIncludes"`
    }
  }
}

type tmpLayoutStruct struct {
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
}

type LayoutContainer struct {
	TemplateRowStart string
	TemplateRowEnd string
	TemplateColumnStart string
	TemplateColumnEnd string
	CssClass string
	DisplayMode string
  JustifySelf string
  AlignSelf string
	Component component_data.ComponentData
}

func (this LayoutContainer) GetHtml() string {
	return fmt.Sprintf(`
    <div class="layout-%v %v"
    style="justify-self:%v;align-self:%v;grid-area:%v/%v/%v/%v">
				%v
		</div>`,
		this.DisplayMode,
    this.CssClass,
    this.JustifySelf,
    this.AlignSelf,
    this.TemplateRowStart,
    this.TemplateColumnStart,
    this.TemplateRowEnd,
    this.TemplateColumnEnd,
    this.Component.GetHtml())
}

func (lc *LayoutContainer) UnmarshalJSONFOO(data []byte) error {
	//Extract LayoutProperties from underlying Component
	var tmp tmpLayoutStruct
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	// Instantiate proper Component
	component, err := UnmarshalCustomValue(data, "componentType", "componentData",
		map[string]reflect.Type{
			"text": reflect.TypeOf(components.TextStruct{}),
		})
	if err != nil {
		return err
	}

  fmt.Println("---------------------")
  fmt.Println(tmp.ComponentData.DisplayMode)
  fmt.Println("---------------------")

	lc.CssClass = tmp.ComponentData.CssClass
	lc.DisplayMode = tmp.ComponentData.DisplayMode
	lc.TemplateRowStart = tmp.ComponentData.TemplateRowStart
	lc.TemplateRowEnd = tmp.ComponentData.TemplateRowEnd
	lc.TemplateColumnStart = tmp.ComponentData.TemplateColumnStart
	lc.TemplateColumnEnd = tmp.ComponentData.TemplateColumnEnd
	lc.JustifySelf = tmp.ComponentData.JustifySelf
	lc.AlignSelf = tmp.ComponentData.AlignSelf
	lc.Component = component

	return nil
}

func UnmarshalCustomValue(data []byte, typeField, resultField string, customTypes map[string]reflect.Type) (component_data.ComponentData, error) {
	m := map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	//fmt.Println(m)
	valueBytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	typeName := m[typeField].(string)
	var newObj component_data.ComponentData
	if ty, found := customTypes[typeName]; found {
		newObj = reflect.New(ty).Interface().(component_data.ComponentData)
		if err = json.Unmarshal(valueBytes, &newObj); err != nil {
			return nil, err
		}
	}
	return newObj, nil
}
