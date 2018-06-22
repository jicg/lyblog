package routers

import (
	"github.com/jicg/lyblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})

	////初始化 case
	//ns := beego.NewNamespace("/case",
	//	beego.NSRouter("/case.html", &controllers.CaseController{}),
	//);
	//
	////注册 namespace
	//beego.AddNamespace(ns)
	//
	//jns := beego.NewNamespace("/jie",
	//	beego.NSRouter("/:page.html", &controllers.JieIndexController{}),
	//);
	//beego.AddNamespace(jns)

	message := beego.NewNamespace("/message",
		beego.NSInclude(&controllers.MessageController{}),
	);
	beego.AddNamespace(message)
	jnu := beego.NewNamespace("/user",
		beego.NSInclude(&controllers.UserController{}),
	);
	beego.Router("/u/:id",&controllers.UserController{},"get:Userhome")
	beego.AddNamespace(jnu)
	api := beego.NewNamespace("/api",
		beego.NSInclude(&controllers.ApiController{}),
	);
	beego.AddNamespace(api)
}
