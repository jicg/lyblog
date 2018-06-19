package models

import (
	"time"
	"math/rand"
)

type User struct {
	Id int
	//昵称
	UserName string
	//邮箱
	Email string
	//密码
	Pass string
	//头像地址
	Avatar string
	//是否认证 例： lyblog 作者
	Approve bool
	//性别 男0 女1
	Sex int
	//城市
	City string
	//签名
	Sign string

	//vip等级
	Rmb string
	//用户角色 管理员1 社区之光2 该号已被封-1
	Auth int
	//飞吻
	Experience int
	//加入时间
	JoinTime time.Time
}

type Verification struct {
	Id    int
	Code  string
	Value string
}

func GetRandVerification() (*Verification, error) {
	count, err := o.QueryTable(&Verification{}).Count()
	if err != nil {
		return nil, err
	}

	if (count <= 0) {
		verification := &Verification{Code: "89-20+1*1", Value: "70"}
		if _, err = o.Insert(verification); err != nil {
			return nil, err
		}
		return verification, nil
	}
	randcnt := rand.Intn(int(count) - 1)
	verification := &Verification{}
	if err = o.QueryTable(&Verification{}).Limit(1, randcnt).One(verification); err != nil {
		return nil, err
	}
	return verification, err
}

func GetUser(email, pass string) (*User, error) {
	user := &User{Email: email, Pass: pass}
	if err := o.Read(user, "Email", "Pass"); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{Email: email}
	if err := o.Read(user, "Email"); err != nil {
		return nil, err
	}
	return user, nil
}

func SaveUser(user *User) (int64, error) {
	return o.Insert(user)
}

func UpdateUser(user *User) (int64, error) {
	return o.Update(user, "UserName", "Email", "Pass", "Avatar", "Approve", "Sex", "City", "Sign")
}
