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

  // gets current date and formats it "Month Day"
  now := time.Now()
  t :=now.Add(-4 * time.Hour)
  time := t.Format("January 2")

  //sets default image and navbar
  this.Data["image"] = luma_response.Payload.Data.ParkingImage.Preview
  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems
  this.Data["data"] = luma_response.Payload.Data

  // loops through all parking options and changes to alternate image if there is a matching date
  for _, element := range luma_response.Payload.Data.Alt {
    if element.ComponentData.AltDate == time {
      this.Data["image"] = element.ComponentData.AltImage.Preview
    }
  }
  
  this.Layout = "layout/layout.tpl"
  this.TplName = "index.tpl"

  this.LayoutSections = make(map[string]string)
  this.LayoutSections["HtmlHead"] = "html_head.tpl"
  this.LayoutSections["HeaderContent"] = "layout/header_content.tpl"
  this.LayoutSections["FooterContent"] = "layout/footer_content.tpl"
  this.LayoutSections["Scripts"] = ""
}