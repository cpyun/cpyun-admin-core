package notify

import (
	"github.com/cpyun/cpyun-admin-core/storage/notify/email"
	"gopkg.in/gomail.v2"
	"testing"
)

const (
	testUser = "user"
	testPwd  = "pwd"
	testHost = "smtp.example.com"
)

const (
	testSSLPort = 465
)

func TestNotify(t *testing.T) {
	client := gomail.NewDialer(testHost, testSSLPort, testUser, testPwd)

	notify := NewNotify()
	notify.Add(
		email.NewEmail(client,
			email.WithDisable(false),
			email.WithFrom("test0@test.com"),
			email.WithTo("test@test.com"),
			email.WithCc("c1@test.com"),
			email.WithBcc("b1@test.com"),
			email.WithBody("text/plain", "hello"),
			email.WithSubject("test"),
			email.WithEmbed("./1.jpg"),
			email.WithAttr("./1.zip"),
		),
	)
}
