package models

import (
	common "github.com/Lumavate-Team/lumavate-go-common"
	widget "github.com/Lumavate-Team/lumavate-go-common/models"
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"fmt"
	"encoding/json"
)

type MainController struct {
  common.LumavateController
}

type TileStruct struct {
	ComponentData struct {
		Title string
	}
}

type LayoutContainer struct {
	TemplateRowStart string
	TemplateRowEnd string
	TemplateColumnStart string
	TemplateColumnEnd string
	Component component_data.ComponentData
}

func (lc *LayoutContainer) UnmarshalJSON(data []byte) error {
    var tmp LayoutContainer
    err := json.Unmarshal(data, &tmp)
    if err != nil {
      return err
    }
    for _, r := range tmp.Component {
        var obj map[string]interface{}
        err := json.Unmarshal(r, &obj)
        if err != nil {
            return err
        }

        objType := ""
				fmt.Println(obj)
        if t, ok := obj["componentType"].(string); ok {
            objType = t
        }

				//Generate new object from reflect
				switch objType {
					case "tile":
						var newObj TileStruct
					case "quote":
						var newObj component_data.QuoteStruct
				}

        err := json.Unmarshal(obj["componentData"], &newObj)
        if err != nil {
            return err
        }
				lc.Component = newObj
    }
    return nil
}

func (this LayoutContainer) GetHtml() string {
	return this.Component.GetHtml()
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
