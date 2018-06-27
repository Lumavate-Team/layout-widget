package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  "encoding/json"
  "widget/models"
	"strings"
  _"github.com/bitly/go-simplejson"
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


  this.Data["data"] = luma_response.Payload.Data

	this.Layout = "layout/layout.tpl"

	mode := this.GetString("mode")

	if strings.ToLower(mode) != "degraded" {
		this.TplName = "index.tpl"
	} else {
		this.TplName = "degraded.tpl"
	}

}
