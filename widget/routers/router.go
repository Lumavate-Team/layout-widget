package routers

import (
  "widget/controllers"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/plugins/cors"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/:ic/:url_ref/:wid", &controllers.MainController{})
    beego.Router("/:ic/:url_ref/discover/health", &controllers.HealthController{})
    beego.Router("/:ic/:url_ref/discover/properties", &controllers.PropertyController{})
    beego.Router("/:ic/:url_ref/discover/components", &controllers.ComponentController{})
    beego.Router("/:ic/:url_ref/:wid/map", &controllers.MainController{})
    beego.InsertFilter(":ic/:url_ref/static/*", beego.BeforeStatic, cors.Allow(&cors.Options{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"GET"},
    AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin"},
    ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
    AllowCredentials: false,
}))

}
