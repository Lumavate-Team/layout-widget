package models

import (
	_"github.com/Lumavate-Team/go-properties/component_data"
  	ims_go_components "github.com/Lumavate-Team/ims-go-components"
    ims_models "github.com/Lumavate-Team/ims-go-components/models"
)

type MainController struct {
  ims_go_components.LumavateController
}

type ContactStruct struct {
	ComponentData struct {
		FirstName string
		LastName string
		JobTitle string `json:"jobTitle"`
		PhoneNumber string
		Email string
	}
}

type LumavateRequest struct {
  Payload struct {
    Data struct {
      ims_models.WidgetStruct
      PrimaryContact []ContactStruct `json:"primaryContacts"`
      SecondaryContact []ContactStruct `json:"secondaryContacts"`
    }
  }
}