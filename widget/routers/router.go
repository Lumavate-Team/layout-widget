package routers

import (
	common_controllers "github.com/Lumavate-Team/lumavate-go-common/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"os"
	"widget/controllers"
)

func init() {
	beego.Router("/:ic/:url_ref/:wid", &controllers.MainController{})
	beego.Router("/:ic/:url_ref/discover/health", &common_controllers.HealthController{})
	beego.Router("/:ic/:url_ref/discover/routes", &controllers.RouteController{})
	beego.Router("/:ic/:url_ref/:wid/single-use-token", &controllers.SingleUseTokenController{})
	beego.Router("/:ic/:url_ref/discover/properties", &controllers.PropertyController{})
	beego.Router("/:ic/:url_ref/instances/:wid/on-create-version", &controllers.VersionCreateController{})
	beego.InsertFilter("/:ic/:url_ref/"+os.Getenv("PUBLIC_KEY")+"/static/*", beego.BeforeStatic, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: false,
	}))
}
