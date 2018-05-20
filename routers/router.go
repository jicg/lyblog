package routers

import (
	"github.com/jicg/lyblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//初始化 case
	ns := beego.NewNamespace("/case",
		beego.NSRouter("/case.html", &controllers.CaseController{}),
	);

	//注册 namespace
	beego.AddNamespace(ns)

	jns := beego.NewNamespace("/jie",
		beego.NSRouter("/:page.html", &controllers.JieIndexController{}),
	);
	beego.AddNamespace(jns)

	jnu := beego.NewNamespace("/user",
		beego.NSRouter("/:page.html", &controllers.UserController{}),
	);
	beego.AddNamespace(jnu)
}
