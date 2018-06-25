package controllers

import (
  "github.com/astaxie/beego"
)

type PropertyController struct {
  beego.Controller
}

func (this *PropertyController) Get() {
  lp := &LumavateProperties{this.Ctx.Request.Header.Get("Authorization"), []*DynamicComponent{}}
  lp.LoadAllComponentSets()
  this.Data["json"] = lp.GetAllProperties()
  this.ServeJSON()
}

