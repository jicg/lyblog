package main

import (
	_ "github.com/jicg/lyblog/routers"
	"github.com/astaxie/beego"
	_ "github.com/jicg/lyblog/models"
	"reflect"
	"fmt"
	"strconv"
	"path"
)

func init() {

}
func main() {
	//sessionon = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "lyblog-key"
	//beego.BConfig.WebConfig.Session.SessionProvider = "file"

	beego.AddFuncMap("eqd", func(x, y interface{}) bool {
		return reflect.DeepEqual(x, y)
	})
	beego.AddFuncMap("eq2", func(x, y interface{}) bool {
		return fmt.Sprintf("%v", x) == fmt.Sprintf("%v", y)
	})
	beego.AddFuncMap("noteq2", func(x, y interface{}) bool {
		return !(fmt.Sprintf("%v", x) == fmt.Sprintf("%v", y))
	})
	beego.AddFuncMap("multiply", func(x, y interface{}) int {
		if x == nil || y == nil {
			return 0
		}
		x1, err := strconv.Atoi(fmt.Sprintf("%v", x))
		if err != nil {
			return 0
		}

		y1, err := strconv.Atoi(fmt.Sprintf("%v", x))
		if err != nil {
			return 0
		}
		return x1 * y1
	})
	beego.SetStaticPath("asset", path.Join("data", "asset"))
	beego.Run()
}
