package integrations

import (
	"fmt"
)

type MailSenderFunc func(from, to, subject, body string) error

func SendMailNotificationForLoginCode(from string, sendMail MailSenderFunc) func(to, message string) error {
	return func(to, message string) error {
		subject := "Mahala - Login code"
		body := fmt.Sprintf("Kod za prijavu: %s", message)

		return sendMail(from, to, subject, body)
	}
}
