package routers

import (
  "widget/controllers"
  "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/:ic/:url_ref/:wid", &controllers.MainController{})
    beego.Router("/:ic/:url_ref/discover/health", &controllers.HealthController{})
    beego.Router("/:ic/:url_ref/discover/properties", &controllers.PropertyController{})
    beego.Router("/:ic/:url_ref/discover/components", &controllers.ComponentController{})
}
