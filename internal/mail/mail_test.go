package mail_test

import (
	"github.com/mahalahub/mahala/internal/mail"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMailSender_Send(t *testing.T) {
	mailU := os.Getenv("MAIL_U")
	mailP := os.Getenv("MAIL_P")

	mailSender := mail.NewServer(mail.Config{
		Host:          "smtp.gmail.com",
		Port:          587,
		Username:      mailU,
		Password:      mailP,
		UseEncryption: true,
	})

	err := mailSender.Send("semirm.dev@gmail.com", "test")
	assert.NoError(t, err)
}
