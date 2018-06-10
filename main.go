package main

import (
	_ "github.com/jicg/lyblog/routers"
	"github.com/astaxie/beego"
	_ "github.com/jicg/lyblog/models"
)

func init() {

}
func main() {
	//sessionon = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "lyblog-key"
	//beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.Run()
}
