package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
  "encoding/json"
  "widget/models"
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

  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["data"] = luma_response.Payload.Data
  this.TplName = "index.tpl"
}
