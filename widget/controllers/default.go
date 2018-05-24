package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  "encoding/json"
  "widget/models"
	"os"
	"fmt"
	"strings"
  "github.com/bitly/go-simplejson"
  "reflect"
  "widget/models/components"
)

type MainController struct {
  common_controller.LumavateController
}

func (this *MainController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  data, err := simplejson.NewJson(this.LumavateGetData())
 
  fmt.Println(data)

  if err != nil {
    this.Abort("500")
  }

  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["formItems"] = luma_response.Payload.Data.FormItems
  fmt.Println(luma_response.Payload.Data.FormItems)

  for _, element := range luma_response.Payload.Data.GridItems {
      if reflect.TypeOf(element.Component) == reflect.TypeOf(components.FormStruct{}) {
        fmt.Println("IN IF STATEMNT")
        element.FormItems.ComponentData.FormInputs.FormItems = luma_response.Payload.Data.FormItems
      }
    }
  


  this.Data["data"] = luma_response.Payload.Data
	this.Data["dnsInfo"] = fmt.Sprintf("%s%s", os.Getenv("PROTO"), this.Ctx.Input.Host())

	this.Layout = "layout/layout.tpl"
	mode := this.GetString("mode")
	if strings.ToLower(mode) != "degraded" {
		this.TplName = "index.tpl"
	} else {
		this.TplName = "degraded.tpl"
	}

  this.LayoutSections["HtmlHead"] = "html_head.tpl"
  this.LayoutSections["FooterContent"] = "home_footer.tpl"
}
