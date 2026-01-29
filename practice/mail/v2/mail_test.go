package v2_test

import (
	"fmt"
	v2 "mail/v2"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

type FakeEmailSender struct {
	subject string
	from    string
	to      []string
	content string
}

type EmailSenderAdapter struct {
	e *email.Email
}

func (s *FakeEmailSender) Send(subject, form string, to []string, content string, mailserver string, a smtp.Auth) error {
	s.subject = subject
	s.from = form
	s.to = to
	s.content = content
	return nil
}

func (adapter *EmailSenderAdapter) Send(subject, from string, to []string, content string, mailserve string, a smtp.Auth) error {
	adapter.e.Subject = subject
	adapter.e.From = from
	adapter.e.To = to
	adapter.e.Text = []byte(content)
	return adapter.e.Send(mailserve, a)
}

func TestSendMailWithDisclaimer2(t *testing.T) {
	s := &FakeEmailSender{}
	err := v2.SendMailWithDisclaimer2(s, "gopher mail test v2", "your_mailbox",
		[]string{"dest_mailbox"},
		"hello, gopher",
		"stmp.163.com:25",
		smtp.PlainAuth("", "your_email_account", "your_email_passwd", "smtp.163.com"))

	if err != nil {
		t.Fatalf("want: nil, actual: %s\n", err)
		return
	}

	want := "hello, gopher" + "\n\n" + v2.DISCLAIMER
	if s.content != want {
		t.Fatalf("want: %s, actual: %s\n", want, s.content)
	}
}

func ExampleSendMailWithDisclaimer2() {
	adapter := &EmailSenderAdapter{
		e: email.NewEmail(),
	}

	err := v2.SendMailWithDisclaimer2(adapter, "gopher mail test v2",
		"your_mailbox",
		[]string{"dest_mailbox"},
		"hello, gopher",
		"stmp.163.com:25",
		smtp.PlainAuth("", "your_email_account", "your_email_passwd", "smtp.163.com"))

	if err != nil {
		fmt.Printf("SendMail error %s\n", err)
		return
	}

	fmt.Println("SendMail success")
}
