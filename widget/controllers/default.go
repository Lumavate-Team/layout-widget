package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  "encoding/json"
  "widget/models"
	"os"
	"fmt"
	"strings"
)

type MainController struct {
  common_controller.LumavateController
}

func (this *MainController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }

  fmt.Println("This is the form branch")

  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
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
