package controllers

import (
  "github.com/astaxie/beego"
)

type ComponentController struct {
  beego.Controller
}

func (this *ComponentController) Get() {
  lp := &LumavateProperties{this.Ctx.Request.Header.Get("Authorization"), []*DynamicComponent{}}
  lp.LoadAllComponentSets()
  this.Data["json"] = lp.GetAllComponents()
  this.ServeJSON()
}
