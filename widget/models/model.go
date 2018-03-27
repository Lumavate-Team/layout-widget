package models

import (
	"github.com/Lumavate-Team/go-properties/component_data"
  	ims_go_components "github.com/Lumavate-Team/ims-go-components"
)

type MainController struct {
  ims_go_components.LumavateController
}

type LumavateRequest struct {
  Payload struct {
    Data struct {
      PageType component_data.PageTypeStruct
      Title ims_go_components.ImsTitleStruct
      ParkingImage component_data.ImageStruct
      NavBarItems ims_go_components.NavBarItemsStruct `json:"navBarItems"`
      NavBar ims_go_components.NavBarStruct `json:"navBar"`
      BackgroundColor string
      AltDate string
      AltImage component_data.ImageStruct
    }
  }
}