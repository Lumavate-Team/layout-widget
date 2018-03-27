package controllers

import (
  ims_go_components "github.com/Lumavate-Team/ims-go-components"
 "fmt"
  "log"
  "time"
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
  this.Data["data"] = luma_response.Payload.Data

  now := time.Now()
  t :=now.Add(-4 * time.Hour)
  time := t.Format("January 2")

  fmt.Println(time)
  fmt.Println(luma_response.Payload.Data.AltDate)

  if luma_response.Payload.Data.AltDate == time {
    this.Data["image"] = luma_response.Payload.Data.AltImage.Preview
    fmt.Println("ALTERNATE IMAGE USED")
  } else {
    this.Data["image"] = luma_response.Payload.Data.ParkingImage.Preview
    fmt.Println("ORIGINAL IMAGE USED")
  }
  
  this.Layout = "layout/layout.tpl"
  this.TplName = "index.tpl"

  this.LayoutSections = make(map[string]string)
  this.LayoutSections["HtmlHead"] = "html_head.tpl"
  this.LayoutSections["HeaderContent"] = "layout/header_content.tpl"
  this.LayoutSections["FooterContent"] = "layout/footer_content.tpl"
  this.LayoutSections["Scripts"] = ""
}