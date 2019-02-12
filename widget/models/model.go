package models

import (
  "fmt"
  common_controllers "github.com/Lumavate-Team/lumavate-go-common/controllers"
  widget "github.com/Lumavate-Team/lumavate-go-common/models"
  component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
)

type MainController struct {
  common_controllers.LumavateController
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

type SecurityOptions struct {
  ComponentHtml string
  ComponentType string
  ComponentData struct {  
    NoAuthRedirect      component_data.PageLinkStruct
    SpecificGroup       []string  `json:"specificGroup"`
  }
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

type AddToHomeStruct struct {
  ShowAddToHome bool `json:"showAddToHome"`
  SkipFirst bool `json:"skipFirst"`
  StartDelay int `json:"startDelay"`
  Lifespan int `json:"lifespan"`
  DisplayCount int `json:"displayCount"`
  Message string `json:"message"`
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
      SecurityProperties     SecurityOptions
      BodyProperties         BodyOptions
      BodyItems              []LayoutContainer
      Footer                 widget.Component
      Header                 widget.Component
      ModalItems             []widget.Component
      HomeScreen             AddToHomeStruct
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
    DisplayMode         string
    JustifySelf         string
    AlignSelf           string
  }
  ComponentHtml string
}

// structs used for getting designer defined user groups
type AuthGroupRequest struct {
  Payload struct {
    Data []GroupStruct
  }
}
type GroupStruct struct {
  Group string `json:"name"`
}

// struct used to get auth-url for making api calls
type AuthRequest struct {
  AuthUrl     string        `json:"authUrl"`
}

// struct used to get login status of user
type GroupRequest struct {  
  Payload struct {
    Data struct{
      Roles       []string `json:"roles"`
      Status      string `json:"status"`
    }
  }
}

func (this LogicContainer) GetHtml() string {
	return fmt.Sprintf(`
	%v
	`,
		this.ComponentHtml)
}

func (this LayoutContainer) GetHtml() string {
  return fmt.Sprintf(`
  <div class="layout-%v"
  style="justify-self:%v;align-self:%v;grid-area:%v/%v/ span %v/ span %v">
  %v
  </div>`,
    this.ComponentData.DisplayMode,
    this.ComponentData.JustifySelf,
    this.ComponentData.AlignSelf,
    this.ComponentData.TemplateRowStart,
    this.ComponentData.TemplateColumnStart,
    this.ComponentData.TemplateRowSpan,
    this.ComponentData.TemplateColumnSpan,
    this.ComponentHtml)
}
