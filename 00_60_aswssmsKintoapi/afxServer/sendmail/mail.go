package sendmail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	UserName   string
	Password   string
	Subject    string
	Body       string
	CC         bool
	BCC        bool
	From       string
	To         string
	HostName   string
	Port       int
	Recipients []string
}

// SMTP送信
func (m *Mail) Send() error {
	//　メールサーバー認証
	auth := smtp.CRAMMD5Auth(m.UserName, m.Password)
	//　メール本文　改行置き換え
	msg := []byte(strings.ReplaceAll(
		fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(m.Recipients, ","), m.Subject, m.Body), "\n", "\r\n"))

	addr := fmt.Sprintf("%s:%d", m.HostName, m.Port)

	if err := smtp.SendMail(addr, auth, m.From, m.Recipients, msg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	//　メールSMTP送信
	mail := &Mail{
		UserName:   "en.b@di-square.co.jp",
		Password:   "6308En",
		Subject:    "メールテスト",
		Body:       "テスト　テスト",
		CC:         false,
		BCC:        false,
		From:       "",
		To:         "en.b@di-square.co.jp",
		HostName:   "smtp.di-square.co.jp",
		Port:       587,
		Recipients: []string{"en.b@di-square.co.jp"},
	}

	//　メール送信
	mail.Send()
}
