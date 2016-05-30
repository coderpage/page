package mail

import "testing"

const (
	mailAddr = "820631390@qq.com"
)

func TestSendMail(t *testing.T) {
	mailer, err := NewServiceMailer()
	if err != nil {
		t.Fatal("NewServiceMailer Err:", err)
	}

	body := `您好！
<br>
<br>
<a href="#">激活链接</a>
<br>
<br>
本邮件为测试邮件，请勿回复！
`
	mailer.SendMail(mailAddr, "uAuth", "Service 邮件测试", "html", body)
}
