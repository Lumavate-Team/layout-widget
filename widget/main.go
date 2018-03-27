package main

import (
  "github.com/Lumavate-Team/go-properties/component_data"
	_ "widget/routers"
	"github.com/astaxie/beego"
  "html/template"
	"os"
)

func ComponentHtml(in component_data.ComponentData) (out template.HTML){
    out = template.HTML(in.GetHtml())
    return
}

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "static","static")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "lc","/lumavate-components/dist")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "ims","/ims-lumavate-components/dist")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "ims-go","/go/src/github.com/Lumavate-Team/ims-go-components/static")
	beego.AddFuncMap("componentHtml", ComponentHtml)
	beego.Run()
}

