package main

import (
	_ "widget/routers"
	models "widget/models"
	"github.com/astaxie/beego"
  "html/template"
	"os"
)

func ComponentHtml(in models.LayoutContainer) (out template.HTML){
		out = template.HTML(in.GetHtml())
    return
}

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "static","static")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "lc","/lumavate-components/dist")
	beego.AddFuncMap("componentHtml", ComponentHtml)
	beego.Run()
}

