package main

import (
	_ "widget/routers"
	models "widget/models"
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"github.com/astaxie/beego"
  "html/template"
	"os"
)

func ComponentHtml(in component_data.ComponentData) (out template.HTML){
		out = template.HTML(in.GetHtml())
    return
}

func LayoutHtml(in models.LayoutContainer) (out template.HTML){
		out = template.HTML(in.GetHtml())
    return
}

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "static","static")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "lc","/lumavate-components/dist")
	beego.AddFuncMap("componentHtml", ComponentHtml)
	beego.AddFuncMap("layoutHtml", LayoutHtml)
	beego.Run()
}

