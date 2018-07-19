package main

import (
	common "github.com/Lumavate-Team/lumavate-go-common"
	"github.com/astaxie/beego"
	"html/template"
	"os"
	models "widget/models"
	_ "widget/routers"
)

func LayoutHtml(in models.LayoutContainer) (out template.HTML) {
	out = template.HTML(in.GetHtml())
	return
}

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX")+os.Getenv("PUBLIC_KEY")+"/static", "static")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX")+"lc", "/node_modules/@lumavate/components/dist")
	beego.AddFuncMap("componentHtml", common.ComponentHtml)
	beego.AddFuncMap("modalHtml", common.ModalHtml)
	beego.AddFuncMap("safeCss", common.SafeCss)
	beego.AddFuncMap("safeHtml", common.SafeHtml)
	beego.AddFuncMap("layoutHtml", LayoutHtml)
	beego.Run()
}
