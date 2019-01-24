package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common/controllers"
  _"github.com/bitly/go-simplejson"
  b64 "encoding/base64"
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
  luma_domain := models.LumavateDomain {}
  token_response := models.AuthRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  // get middle of pwa_jwt token
  token := this.Ctx.GetCookie("pwa_jwt")
  if (token != "") {
    token = strings.Split(token, ".")[1]

    // add padding to jwt if number of bytes is not correct
    if i := len(token) % 4; i != 0 {
      token += strings.Repeat("=", 4-i)
    }

    // decode the token and ummarshal into auth struct
    decodedToken, _ := b64.StdEncoding.DecodeString(token)
    if err := json.Unmarshal(decodedToken, &token_response); err != nil {
      panic(err)
    }

    // call out to get login status of user
    q := fmt.Sprintf("%vstatus",token_response.AuthUrl)
    fmt.Println(q)
    user_roles := models.GroupRequest {}
    groups, status := this.LumavateGet(q, true)
    if err := json.Unmarshal(groups, &user_roles); err != nil {
      fmt.Println(err)
    }

    // if user does not have access send them to login page
    if luma_response.Payload.Data.SecurityProperties.ComponentType != "securityNone"  {
      if (status == "401") {
        this.Ctx.Redirect(302, token_response.AuthUrl + "login")
      }
    }

    // check if user is allowed to access specific page
    // if not redirect them somewhere
    if luma_response.Payload.Data.SecurityProperties.ComponentType == "securitySpecific" {
      if (this.checkRole(user_roles.Payload.Data.Roles, luma_response.Payload.Data.SecurityProperties.ComponentData.SpecificGroup) != true) {
        this.Ctx.Redirect(302, luma_response.Payload.Data.SecurityProperties.ComponentData.NoAuthRedirect.Url)
      } 
    }
  }

  // all auth has been dealt with continue normal page loading if qualified
  domain, _ := this.LumavateGet("/pwa/v1/domain")
  json.Unmarshal(domain, &luma_domain)

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
  this.Data["gtm"] = luma_domain.Payload.Data.RuntimeData["gtm"]

  // fmt.Println(this.XSRFToken())

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








