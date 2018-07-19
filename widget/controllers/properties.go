package controllers

import (
	"github.com/Lumavate-Team/lumavate-go-common/properties"
	"github.com/astaxie/beego"
)

type PropertyController struct {
	beego.Controller
}

func (this *PropertyController) Get() {
	lp := WidgetProperties{properties.NewLumavateProperties(this.Ctx.Request.Header.Get("Authorization"))}
	this.Data["json"] = lp.GetAllProperties()
	this.ServeJSON()
}
