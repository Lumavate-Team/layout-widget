package models

import (
  common_models "github.com/Lumavate-Team/lumavate-go-common/models"
  "html/template"
)

type LumavateRequest struct {
  InstanceName      string                        `json:"instance__name"`
  PageType          common_models.PageTypeStruct  `json:"pageType"`
  DirectIncludes    [] string                     `json:"__directIncludes"`
  DirectCssIncludes [] string                     `json:"__directCssIncludes"`
  StyleData         [] struct {
    Name string
    Value string
  }
  BackgroundColor        string
  DisplayBackgroundImage bool
  BackgroundImage        common_models.ImageStruct
  DisplayHeader          bool
  DisplayFooter          bool
  SecurityProperties  struct {
    common_models.ComponentStruct
    ComponentData struct {
      NoAuthRedirect      common_models.PageLinkStruct
      SpecificGroup       [] string  `json:"specificGroup"`
    }
  }
  Script template.JS
  ViewModel template.JS
  ViewTemplate template.HTML
  BodyProperties         struct {
    common_models.ComponentStruct
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
			PaddingTop          int
			PaddingRight        int
			PaddingLeft         int
			PaddingBottom       int
    }
  }
  Variables [] struct {
		ComponentType string
    ComponentData struct {
      VariableId          template.HTML
			StringValue         template.HTML
			IntValue            int
			ColorValue          template.HTML
		  ImageValue struct {
				Preview string
				PreviewSmall string
				PreviewMedium string
				PreviewLarge string
			}
    }
  }
  Translations [] struct {
    ComponentData struct {
      StringId            template.HTML
      String              template.HTML
    }
  }
  Templates [] struct {
    ComponentData struct {
      TemplateId          template.HTML
      Template            template.HTML
    }
  }
  BodyItems              [] LayoutContainerStruct
  Footer                 common_models.ComponentStruct
  Header                 common_models.ComponentStruct
  ModalItems             [] common_models.ComponentStruct
  HomeScreen              struct {
    ShowAddToHome bool
    SkipFirst bool
    StartDelay int
    Lifespan int
    DisplayCount int
    Message string
  }
  LogicItems []struct {
    common_models.ComponentStruct
    ComponentData struct {
      Placement string
    }
  }
}
