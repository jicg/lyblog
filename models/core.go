package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var (
	o orm.Ormer
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Verification), new(Note),new(Replay))
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RunSyncdb("default", false, true)
	 orm.Debug = true
	o = orm.NewOrm()
	o.Using("default")
	count, _ := o.QueryTable(&User{}).Count();
	if count == 0 {
		o.Insert(&User{UserName: "admin",
			//邮箱
			Email: "admin@qq.com",
			//密码
			Pass: "123123",
			//头像地址
			Avatar: "/static/images/avatar/default.png",
			//是否认证 例： lyblog 作者
			Approve: "lyblog 作者",
			//性别 男0 女1
			Sex: 0,
			//城市
			City: "淮安",
			//签名
			Sign: "坚持就是胜利",

			//vip等级
			Rmb: "无敌vip",
			//用户角色 管理员1 社区之光2 该号已被封-1
			Auth: 1,
			//飞吻
			Experience: 0,
			//加入时间
			JoinTime: time.Now(),
		})
	}
}
