package main

import (
  common "github.com/Lumavate-Team/lumavate-go-common"
  "github.com/astaxie/beego"
  "html/template"
  "os"
  models "widget/models"
  _ "widget/routers"
  "strings"
)

func LayoutHtml(in models.LayoutContainerStruct) (out template.HTML) {
  out = template.HTML(in.GetHtml())
  return
}

func LogicHtml(in models.LogicContainerStruct) (out template.HTML){
  out = template.HTML(in.GetHtml())
  return
}

func HasSuffix(in string, test string) (out bool){
  return strings.HasSuffix(in, test)
}
func HasPrefix(in string, test string) (out bool){
  return strings.HasPrefix(in, test)
}
func Replace(input, from,to string) string {
	return strings.Replace(input,from,to, -1)
}
func main() {
  beego.BConfig.WebConfig.EnableXSRF = false
  beego.BConfig.WebConfig.XSRFKey = os.Getenv("PRIVATE_KEY")
  beego.BConfig.WebConfig.XSRFExpire = 3600

  beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX")+os.Getenv("PUBLIC_KEY")+"/static", "static")
  beego.SetStaticPath(os.Getenv("WIDGET_URL_PREFIX")+os.Getenv("PUBLIC_KEY")+"/core", "/lumavate-core-components")
  beego.AddFuncMap("componentHtml", common.ComponentHtml)
  beego.AddFuncMap("modalHtml", common.ModalHtml)
  beego.AddFuncMap("safeCss", common.SafeCss)
  beego.AddFuncMap("safeHtml", common.SafeHtml)
  beego.AddFuncMap("layoutHtml", LayoutHtml)
  beego.AddFuncMap("logicHtml", LogicHtml)
  beego.AddFuncMap("hasSuffix", HasSuffix)
  beego.AddFuncMap("hasPrefix", HasPrefix)
  beego.AddFuncMap("replace", Replace)
  beego.Run()
}
