package models

import (
	common "github.com/Lumavate-Team/lumavate-go-common"
	widget "github.com/Lumavate-Team/lumavate-go-common/models"
	_"github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type MainController struct {
  common.LumavateController
}

type TileStruct struct {
	ComponentData struct {
		Title string
		TemplateRowStart string
		TemplateRowEnd string
		TemplateColumnStart string
		TemplateColumnEnd string
	}
}

type LumavateRequest struct {
  Payload struct {
    Data struct {
      widget.CommonWidgetStruct
			Padding int
			GridTemplateColumns string
			GridTemplateRows string
			Tiles []TileStruct `json:"tiles"`
    }
  }
}
