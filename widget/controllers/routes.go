package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common/controllers"
)

type RouteStruct struct {
		Path string `json:"path"`
		Security [] string `json:"security"`
		Type string `json:"type"`
}

type RouteController struct {
  common_controller.LumavateController
}

func (this *RouteController) Get() {
    this.Ctx.Output.SetStatus(200)
    result := RouteStruct{"*", []string {"jwt"}, "page"}
		routes := [] RouteStruct {result}
    this.Data["json"] = routes
    this.ServeJSON()
}
