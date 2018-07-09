package controllers

import (
  "github.com/astaxie/beego"
	"fmt"
)

type SingleUseTokenController struct {
  beego.Controller
}

func (this *SingleUseTokenController) Post() {
	fmt.Println(this.Ctx.Input.Header("Authorization"))

  result := HealthStruct{ "Ok" }
  this.Data["json"] = &result
  this.ServeJSON()
}
