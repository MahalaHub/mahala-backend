package integrations

import (
	mail "github.com/xhit/go-simple-mail"
	"time"
)

type MailSender struct {
	smtpServer *mail.SMTPServer
}

type MailSenderConfig struct {
	Host          string
	Port          int
	Username      string
	Password      string
	UseEncryption bool
}

func NewMailSender(conf MailSenderConfig) MailSender {
	server := mail.NewSMTPClient()

	server.Host = conf.Host
	server.Port = conf.Port
	server.Username = conf.Username
	server.Password = conf.Password
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	if conf.UseEncryption {
		server.Encryption = mail.EncryptionTLS
	}

	return MailSender{
		smtpServer: server,
	}
}

func (m MailSender) Send(to, message string) error {
	var from, cc, bcc, subject string

	return m.send(from, to, cc, bcc, subject, message)
}

func (m MailSender) send(from, to, cc, bcc, subject, body string) error {
	smtpClient, err := m.smtpServer.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()

	email.SetFrom(from)
	email.AddTo(to)
	email.AddCc(cc)
	email.AddBcc(bcc)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	return email.Send(smtpClient)
}
