package main

import (
	_ "widget/routers"
	"github.com/astaxie/beego"
  "html/template"
	"os"
	"fmt"
)

func ComponentHtml(in map[string]interface{}) (out template.HTML){
	fmt.Println(in)
		out = "Hello"
    return
}

func main() {
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "static","static")
	beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "lc","/lumavate-components/dist")
	beego.AddFuncMap("componentHtml", ComponentHtml)
	beego.Run()
}

