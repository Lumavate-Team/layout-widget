package controllers

import (
  "github.com/astaxie/beego"
)

type ComponentController struct {
  beego.Controller
}

func (this *ComponentController) Get() {
  lp := &LumavateProperties{}
  this.Data["json"] = lp.GetAllComponents()
  this.ServeJSON()
}
