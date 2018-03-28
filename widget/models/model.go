package models

import (
	"github.com/Lumavate-Team/go-properties/component_data"
  	ims_go_components "github.com/Lumavate-Team/ims-go-components"
    ims_models "github.com/Lumavate-Team/ims-go-components/models"
)

type MainController struct {
  ims_go_components.LumavateController
}

type AltParkingStruct struct {
  ComponentData struct {
    AltDate string
    AltImage component_data.ImageStruct
  }
}

type LumavateRequest struct {
  Payload struct {
    Data struct {
      ims_models.WidgetStruct
      ParkingImage component_data.ImageStruct
      Alt []AltParkingStruct `json:"alternateParking"`
    }
  }
}