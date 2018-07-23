package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  "encoding/json"
  "widget/models"
  "strings"
	"fmt"
)

type MainController struct {
  common_controller.LumavateController
}

func (this *MainController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

	if luma_response.Payload.Data.BodyProperties.ComponentType == "body-items-basic" {
		body_props := &luma_response.Payload.Data.BodyProperties.ComponentData
		body_props.BodyTemplateRows = fmt.Sprintf("repeat(%v, 1fr)", body_props.BodyNumRows)
		body_props.BodyTemplateColumns = fmt.Sprintf("repeat(%v, 1fr)", body_props.BodyNumColumns)
		if body_props.BodyMaxWidth != 0 {
			body_props.BodyMaxWidthStr = fmt.Sprintf("%vpx", body_props.BodyMaxWidth)
		} else {
			body_props.BodyMaxWidthStr = "100%"
		}
	}

  if err != nil {
    this.Abort("500")
  }

  this.Data["data"] = luma_response.Payload.Data
  this.Layout = "layout/layout.tpl"

  mode := this.GetString("mode")

  if strings.ToLower(mode) != "degraded" {
    this.TplName = "index.tpl"
    this.Data["degraded"] = false
  } else {
    this.Data["degraded"] = true
    this.TplName = "degraded.tpl"
  }

}
