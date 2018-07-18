package controllers

import (
  "github.com/astaxie/beego"
)

type PropertyController struct {
  beego.Controller
}

func (this *PropertyController) Get() {
	lp := NewLumavateProperties(this.Ctx.Request.Header.Get("Authorization"))
  this.Data["json"] = lp.GetAllProperties()
  this.ServeJSON()
}

