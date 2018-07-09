package main

import (
  _ "widget/routers"
  models "widget/models"
  component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
  "github.com/astaxie/beego"
  "html/template"
  "os"
)

func SafeCss(in string) (out template.CSS){
  out = template.CSS(in)
  return
}

func SafeHtml(in string) (out template.HTML){
  out = template.HTML(in)
  return
}

func ComponentHtml(in component_data.ComponentData) (out template.HTML){
  out = template.HTML(in.GetHtml())
  return
}

func LayoutHtml(in models.LayoutContainer) (out template.HTML){
  out = template.HTML(in.GetHtml())
  return
}

func main() {
	beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFKey = os.Getenv("PRIVATE_KEY")
	beego.BConfig.WebConfig.XSRFExpire = 3600

  beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + os.Getenv("PUBLIC_KEY") + "/static","static")
  beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX") + "lc","/lumavate-components/dist")
  beego.AddFuncMap("componentHtml", ComponentHtml)
  beego.AddFuncMap("layoutHtml", LayoutHtml)
  beego.AddFuncMap("safeCss", SafeCss)
  beego.AddFuncMap("safeHtml", SafeHtml)
  beego.Run()
}

