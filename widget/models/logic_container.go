package models

import (
	"fmt"
  common_models "github.com/Lumavate-Team/lumavate-go-common/models"
)

type LogicContainerStruct struct {
	common_models.ComponentStruct
	ComponentData struct {
		Placement string	
	}
}

func (this LogicContainerStruct) GetHtml() string {
	return fmt.Sprintf(`
	%v
	`,
		this.ComponentHtml)
}

