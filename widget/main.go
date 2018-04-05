package main

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	_ "widget/routers"
	"github.com/astaxie/beego"
  "html/template"
	"os"
	"fmt"
)

func ComponentHtml(in component_data.ComponentData) (out template.HTML){
    out = template.HTML(in.GetHtml())
    return
}

func SafeHTML(in string) (out template.HTML){
    out = template.HTML(in)
		fmt.Println("SafeHTML")
		fmt.Println(out)
    return
}

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "static","static")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "lc","/lumavate-components/dist")
	beego.AddFuncMap("componentHtml", ComponentHtml)
	beego.AddFuncMap("safeHTML", SafeHTML)
	beego.Run()
}

