package controllers

import (
	"encoding/json"
	"fmt"
	common_controller "github.com/Lumavate-Team/lumavate-go-common/controllers"
	_ "html/template"
	_ "os"
	_ "strings"
	"widget/models"
)

type MainController struct {
	common_controller.LumavateController
}

func (this *MainController) Get() {
	response := this.LumavateGetData()

	luma_response := models.LumavateRequest{}
	json.Unmarshal(response.Payload.Data.WidgetData, &luma_response)

	// if user does not have access send them to login page
	if luma_response.SecurityProperties.ComponentType != "securityNone" {
		if response.Payload.Data.AuthData.Status == "inactive" {
			this.Ctx.Redirect(302, response.Payload.Data.TokenData.AuthUrl+"login")
		}
	}

	// check if user is allowed to access specific page
	// if not redirect them somewhere
	if luma_response.SecurityProperties.ComponentType == "securitySpecific" {
		if this.checkRole(response.Payload.Data.AuthData.Roles, luma_response.SecurityProperties.ComponentData.SpecificGroup) != true {
			this.Ctx.Redirect(302, luma_response.SecurityProperties.ComponentData.NoAuthRedirect.Url)
		}
	}

	this.Data["json"] = &luma_response
	//  this.Ctx.Output.SetStatus(200)
	this.ServeJSON()
}

func (this *MainController) checkRole(userRoles []string, definedRoles []string) bool {
	fmt.Println(userRoles)
	fmt.Println(definedRoles)
	for _, user := range userRoles {
		for _, defined := range definedRoles {
			if user == defined {
				return true
			}
		}
	}
	return false
}
