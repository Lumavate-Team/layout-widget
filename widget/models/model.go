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

type LumavateRequest struct {
  Payload struct {
    Data struct {
      widget.CommonWidgetStruct
			Padding int
			GridTemplateColumns string
			GridTemplateRows string
			GridItems []LayoutContainer
    }
  }
}

type tmpLayoutStruct struct {
	ComponentData struct {
		TemplateRowStart string
		TemplateRowEnd string
		TemplateColumnStart string
		TemplateColumnEnd string
	}
}

type LayoutContainer struct {
	TemplateRowStart string
	TemplateRowEnd string
	TemplateColumnStart string
	TemplateColumnEnd string
	Component component_data.ComponentData
}

func (this LayoutContainer) GetHtml() string {
	fmt.Println("Call Layout GetHtml")
	fmt.Println(this.Component)
	return fmt.Sprintf(`
    <div class="grid-item"
		style="position:relative;text-align:center;padding:2px;grid-area:%v/%v/%v/%v">
				%v
		</div>`,
    this.TemplateRowStart,
    this.TemplateColumnStart,
    this.TemplateRowEnd,
    this.TemplateColumnEnd,
    this.Component.GetHtml())
}

func (lc *LayoutContainer) UnmarshalJSON(data []byte) error {
	//Extract LayoutProperties from underlying Component
	var tmp tmpLayoutStruct
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	// Instantiate proper Component
	component, err := UnmarshalCustomValue(data, "componentType", "componentData",
		map[string]reflect.Type{
			"navigation": reflect.TypeOf(components.NavigationStruct{}),
			"video": reflect.TypeOf(components.VideoStruct{}),
			"quote": reflect.TypeOf(component_data.QuoteStruct{}),
		})
	if err != nil {
		return err
	}

	lc.TemplateRowStart = tmp.ComponentData.TemplateRowStart
	lc.TemplateRowEnd = tmp.ComponentData.TemplateRowEnd
	lc.TemplateColumnStart = tmp.ComponentData.TemplateColumnStart
	lc.TemplateColumnEnd = tmp.ComponentData.TemplateColumnEnd
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
	switch typeName {
		case "navigation":
			var newObj components.NavigationStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		case "video":
			var newObj components.VideoStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		case "quote":
			var newObj component_data.QuoteStruct
			if err = json.Unmarshal(valueBytes, &newObj); err != nil {
				return nil, err
			}
			return newObj, nil
		}
	//var newObj component_data.ComponentData
	//if ty, found := customTypes[typeName]; found {
	//	newObj = reflect.New(ty).Interface().(component_data.ComponentData)
	//}
	return nil, nil
}
