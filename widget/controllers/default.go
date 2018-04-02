package controllers

import (
  ims_go_components "github.com/Lumavate-Team/ims-go-components"
 "fmt"
  "log"
  _"time"
  "encoding/json"
  "github.com/bitly/go-simplejson"
  "widget/models"
)

type MainController struct {
  ims_go_components.LumavateController
}

func (this *MainController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  data, err := simplejson.NewJson(this.LumavateGetData())

  fmt.Println(data)
  if err != nil {
    fmt.Println(err)
    log.Fatal(err)
    this.Abort("500")
  }

  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["data"] = luma_response.Payload.Data
  this.Data["primary"] = luma_response.Payload.Data.PrimaryContact
  this.Data["secondary"] = luma_response.Payload.Data.SecondaryContact
  
  this.Layout = "layout/layout.tpl"
  this.TplName = "index.tpl"

  this.LayoutSections = make(map[string]string)
  this.LayoutSections["HtmlHead"] = "html_head.tpl"
  this.LayoutSections["HeaderContent"] = "layout/header_content.tpl"
  this.LayoutSections["FooterContent"] = "layout/footer_content.tpl"
  this.LayoutSections["Scripts"] = ""
}