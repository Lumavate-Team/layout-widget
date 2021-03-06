package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common/controllers"
  "encoding/json"
  "widget/models"
  "strings"
  "fmt"
  "html/template"
  "os"
)

type MainController struct {
  common_controller.LumavateController
}

func (this *MainController) Get() {
  response := this.LumavateGetData()

  luma_response := models.LumavateRequest {}
  json.Unmarshal(response.Payload.Data.WidgetData, &luma_response)

  // if user does not have access send them to login page
  if luma_response.SecurityProperties.ComponentType != "securityNone"  {
    if (response.Payload.Data.AuthData.Status == "inactive") {
      this.Ctx.Redirect(302, response.Payload.Data.TokenData.AuthUrl + "login")
    }
  }

  // check if user is allowed to access specific page
  // if not redirect them somewhere
  if luma_response.SecurityProperties.ComponentType == "securitySpecific" {
    if (this.checkRole(response.Payload.Data.AuthData.Roles, luma_response.SecurityProperties.ComponentData.SpecificGroup) != true) {
      this.Ctx.Redirect(302, luma_response.SecurityProperties.ComponentData.NoAuthRedirect.Url)
    }
  }

  if luma_response.BodyProperties.ComponentType == "body-items-basic" {
    body_props := &luma_response.BodyProperties.ComponentData
    body_props.BodyTemplateRows = fmt.Sprintf("repeat(%v, 1fr)", body_props.BodyNumRows)
    body_props.BodyTemplateColumns = fmt.Sprintf("repeat(%v, 1fr)", body_props.BodyNumColumns)
    if body_props.BodyMaxWidth != 0 {
      body_props.BodyMaxWidthStr = fmt.Sprintf("%vpx", body_props.BodyMaxWidth)
    } else {
      body_props.BodyMaxWidthStr = "100%"
    }
  }

  for i, element := range luma_response.StyleData {
    if strings.HasSuffix(element.Name, "FontFamily") && !(strings.HasPrefix(element.Value, "custom:") || strings.HasPrefix(element.Value, "standard:") || strings.HasPrefix(element.Value, "google:")) {
      luma_response.StyleData[i].Value = "google:" + element.Value
    }
  }

  this.Data["baseUrl"] = fmt.Sprintf("%s%s/",
      os.Getenv("WIDGET_URL_PREFIX"),
      this.Ctx.Input.Param(":wid"),
  )

  this.Data["data"] = luma_response
  this.Data["mode"] = os.Getenv("MODE")
  this.Data["resources"] = response.Payload.Data.Resources
  this.Data["gtm"] = response.Payload.Data.DomainData.RuntimeData["gtm"]
  auth_json, _ := json.Marshal(response.Payload.Data.AuthData)
  this.Data["auth_json"] = template.JS(string(auth_json))
  activation_json, _ := json.Marshal(response.Payload.Data.ActivationData)
  this.Data["activation_json"] = template.JS(string(activation_json))
  domain_json, _ := json.Marshal(response.Payload.Data.DomainData)
  this.Data["domain_json"] = template.JS(string(domain_json))

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

func (this *MainController) checkRole(userRoles []string, definedRoles []string) bool {
  fmt.Println(userRoles)
  fmt.Println(definedRoles)
  for _, user := range userRoles {
    for _, defined := range definedRoles {
      if (user == defined) {
        return true
      }
    }
  }
  return false
}








