package sysutils

import (
	"net/smtp"
	"strings"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"github.com/murlokswarm/errors"
)

var (
	HOST        = beego.AppConfig.String("email::host")
	SERVER_ADDR = beego.AppConfig.String("email::server_addr")
	USER        = beego.AppConfig.String("email::user")     //发送邮件的邮箱
	PASSWORD    = beego.AppConfig.String("email::password") //发送邮件邮箱的密码
)

type Email struct {
	to       string "to"
	subject  string "subject"
	msg      string "msg"
	mailtype string "mailtype"
}

func NewEmail(to, subject, msg string, mailtype string) *Email {
	return &Email{to: to, subject: subject, msg: msg, mailtype: mailtype}
}

func SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	sendTo := strings.Split(email.to, ";")
	var content_type string
	if email.mailtype == "html" {
		content_type = "Content-Type: text/html;charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain;charset=UTF-8"
	}

	done := make(chan error, 1024)

	go func() {
		defer close(done)
		for _, v := range sendTo {
			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~"+content_type+"~~", "~", "\r\n", -1) + email.msg
			err := smtp.SendMail(
				SERVER_ADDR,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)
			done <- err
		}
	}()
	var err1 error
	for i := 0; i < len(sendTo); i++ {
		err := <-done
		if err != nil {
			err1 = errors.New(ConvertToString(err.Error(), "gbk", "utf8"))
			beego.Error(err)
		}
	}

	return err1
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
