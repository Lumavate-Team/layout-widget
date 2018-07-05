package main

import (
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"github.com/astaxie/beego"
	"html/template"
	"os"
	models "widget/models"
	_ "widget/routers"
)

func SafeCss(in string) (out template.CSS) {
	out = template.CSS(in)
	return
}

func SafeHtml(in string) (out template.HTML) {
	out = template.HTML(in)
	return
}

func ComponentHtml(in component_data.ComponentData) (out template.HTML) {
	out = template.HTML(in.GetHtml())
	return
}

func LayoutHtml(in models.LayoutContainer) (out template.HTML) {
	out = template.HTML(in.GetHtml())
	return
}

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX")+os.Getenv("PUBLIC_KEY")+"/static", "static")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX")+"lc", "/node_modules/@lumavate/components/dist")
	beego.AddFuncMap("componentHtml", ComponentHtml)
	beego.AddFuncMap("layoutHtml", LayoutHtml)
	beego.AddFuncMap("safeCss", SafeCss)
	beego.AddFuncMap("safeHtml", SafeHtml)
	beego.Run()
}
