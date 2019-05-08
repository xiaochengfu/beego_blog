package routers

import (
	"github.com/astaxie/beego"
	"github.com/xiaochengfu/beego_blog/controllers"
)

func init() {

	beego.Router("/", &controllers.BlogController{}, "*:Home")
	beego.Router("/home", &controllers.BlogController{}, "*:Home")
	beego.Router("/article", &controllers.BlogController{}, "*:Article")
	beego.Router("/detail", &controllers.BlogController{}, "*:Detail")
	beego.Router("/about", &controllers.BlogController{}, "*:About")
	beego.Router("/timeline", &controllers.BlogController{}, "*:Timeline")
	beego.Router("/resource", &controllers.BlogController{}, "*:Resource")
	beego.Router("/comment", &controllers.BlogController{}, "post:Comment")

	//api
	beego.Router("/api/homeapi", &controllers.BlogController{}, "*:Homeapi")

	//后台admin
	beego.AutoRouter(&controllers.AdminController{})

	//chan学习路由
	beego.Router("/api/run", &controllers.Chan1{}, "*:Run")

}
