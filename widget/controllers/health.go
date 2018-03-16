package controllers

import (
  "github.com/astaxie/beego"
)

type HealthStruct struct {
  Status string `json:"status"`
}

type HealthController struct {
  beego.Controller
}

func (this *HealthController) Get() {
  result := HealthStruct{ "Ok" }
  this.Data["json"] = &result
  this.ServeJSON()
}
