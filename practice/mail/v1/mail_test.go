package mail_test

import (
	"mail/v1"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := mail.SendMailWithDisclaimer("gopher mail test v1", "your_mailbox",
		[]string{"dest_mailbox"},
		"hello, gopher",
		"stmp.163.com:25",
		smtp.PlainAuth("", "your_email_account", "your_email_passwd", "smtp.163.com"))

	if err != nil {
		t.Fatalf("want: nil, actual: %s\n", err)
	}
}
