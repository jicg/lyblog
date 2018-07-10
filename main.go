package main

import (
	"github.com/astaxie/beego"
	_ "github.com/jicg/lyblog/routers"
	_ "github.com/jicg/lyblog/models"
	_ "github.com/jicg/lyblog/controllers"
	"reflect"
	"fmt"
	"strconv"
	"path"
	"os"
	"github.com/astaxie/beego/logs"
	"github.com/jicg/lyblog/sysutils"
	"time"
	"encoding/gob"
	"github.com/jicg/lyblog/models"
)

func main() {
	//text()
	initLog()
	initSession()
	initTemplate()
	beego.SetStaticPath("asset", path.Join("data", "asset"))
	beego.Run()
}

func initLog() {
	if err := os.MkdirAll("data/logs", 0777); err != nil {
		beego.Error(err)
		return
	}
	logs.SetLogger("file", `{"filename":"data/logs/lyblog.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Async(1e3)
}

func initSession() {
	gob.Register(&models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "lyblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate() {
	beego.AddFuncMap("eqd", func(x, y interface{}) bool {
		return reflect.DeepEqual(x, y)
	})
	beego.AddFuncMap("eq2", func(x, y interface{}) bool {
		return fmt.Sprintf("%v", x) == fmt.Sprintf("%v", y)
	})
	beego.AddFuncMap("eg", func(x, y int) bool {
		return x > y
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

		y1, err := strconv.Atoi(fmt.Sprintf("%v", y))
		if err != nil {
			return 0
		}
		beego.Info("multiply ,", x1, ",", y1)
		return x1 * y1
	})
	beego.AddFuncMap("nvl", func(str, def string) string {
		if len(str) == 0 {
			return def
		}
		return str
	})
}

func text() {
	//	email_tpl := `
	//	<table border="0" cellpadding="0" cellspacing="0"
	//       style="width: 600px; border: 1px solid #ddd; border-radius: 3px; color: #555; font-family: 'Helvetica Neue Regular',Helvetica,Arial,Tahoma,'Microsoft YaHei','San Francisco','微软雅黑','Hiragino Sans GB',STHeitiSC-Light; font-size: 12px; height: auto; margin: auto; overflow: hidden; text-align: left; word-break: break-all; word-wrap: break-word;">
	//    <tbody style="margin: 0; padding: 0;">
	//    <tr style="background-color: #393D49; height: 60px; margin: 0; padding: 0;">
	//        <td style="margin: 0; padding: 0;">
	//            <div style="color: #5EB576; margin: 0; margin-left: 30px; padding: 0;">
	//                <a style="font-size: 14px; margin: 0; padding: 0; color: #5EB576; text-decoration: none;"
	//                   href="{{ .Host }}" target="_blank">
	//                    上班了吗？</a>
	//            </div>
	//        </td>
	//    </tr>
	//    <tr style="margin: 0; padding: 0;">
	//        <td style="margin: 0; padding: 30px;">
	//            <p style="line-height: 20px; margin: 0; margin-bottom: 10px; padding: 0;">
	//                你好，<em style="font-weight: 700;">小丽子</em>童鞋，请在30分钟内重置您的密码： </p>
	//            <div style=""><a
	//                    href="{{.AuthUrl}}"
	//                    style="background-color: #009E94; color: #fff; display: inline-block; height: 32px; line-height: 32px; margin: 0 15px 0 0; padding: 0 15px; text-decoration: none;"
	//                    target="_blank">立即重置密码</a></div>
	//            <p style="line-height: 20px; margin-top: 20px; padding: 10px; background-color: #f2f2f2; font-size: 12px;">
	//                如果该邮件不是由你本人操作，请勿进行激活！否则你的邮箱将会被他人绑定。 </p></td>
	//    </tr>
	//    <tr style="background-color: #fafafa; color: #999; height: 35px; margin: 0; padding: 0; text-align: center;">
	//        <td style="margin: 0; padding: 0;">系统邮件，请勿直接回复。</td>
	//    </tr>
	//    </tbody>
	//</table>
	//`
	time.AfterFunc(3000, func() { //3475404009@qq.com
		//sysutils.SendEmail(sysutils.NewEmail("3475404009@qq.com", "text", email_tpl, "html"))
		beego.Info(sysutils.Token())
		text()
	})
}
