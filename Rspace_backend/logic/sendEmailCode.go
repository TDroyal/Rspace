package logic

// 写一个发送验证码的邮箱地址的逻辑

import (
	"crypto/tls"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
)

// 生成随机的验证码
func GenerateRandCode() string {
	//随机数种子设置
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(random.Intn(10))
	}
	return code
}

// 发送验证码
func SendCodeByEmail(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "royal_111 <royal_3@qq.com>" // 发送者
	e.To = []string{toUserEmail}          //接收者
	e.Subject = "Welcome to Rspace"       //主题
	// e.Text = []byte("Text Body is, of course, supported!") //发送文本
	e.HTML = []byte("您的验证码是<b>" + code + "</b>") //发送html
	// if err := e.Send("smtp.qq.com:465", smtp.PlainAuth("", "royal_3@qq.com", "pyzlylyfajftbeeh", "smtp.qq.com:465")); err != nil {
	// 	return err  //报错//EOF
	// }
	if err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "royal_3@qq.com", "pyzlylyfajftbeeh", "smtp.qq.com:465"), &tls.Config{ServerName: "smtp.qq.com:465", InsecureSkipVerify: true}); err != nil { //InsecureSkipVerify: true跳过验证
		return err
	}
	return nil
}
