package mail

import (
	mail "github.com/xhit/go-simple-mail"
	"time"
)

type Server struct {
	smtpServer *mail.SMTPServer
}

type Config struct {
	Host          string
	Port          int
	Username      string
	Password      string
	UseEncryption bool
}

func NewServer(conf Config) Server {
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

	return Server{
		smtpServer: server,
	}
}

func (m Server) Send(from, to, subject, body string) error {
	smtpClient, err := m.smtpServer.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()

	email.SetFrom(from)
	email.AddTo(to)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, body)

	return email.Send(smtpClient)
}
