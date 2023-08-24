package user

import (
	"errors"
	"github.com/mahalahub/mahala/internal/random"
	"github.com/sirupsen/logrus"
)

var existsErr = errors.New("username already exists")

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AccountManagement struct {
	repo             Repository
	sendNotification NotificationSenderFunc
}

type Repository interface {
	GetUserAccount(username string) (Account, error)
	AddUserAccount(account Account) (Account, error)
}

type NotificationSenderFunc func(to, message string) error

func NewAccountManagement(repo Repository, sendNotification NotificationSenderFunc) AccountManagement {
	return AccountManagement{
		repo:             repo,
		sendNotification: sendNotification,
	}
}

func (accMan AccountManagement) GenerateLoginCode(username, email string) (string, error) {
	acc, err := accMan.repo.GetUserAccount(username)
	if err != nil {
		return "", err
	}

	if acc.ID > 0 && acc.Email != email {
		return "", existsErr
	}

	_, err = accMan.repo.AddUserAccount(Account{
		Username: username,
		Email:    email,
	})
	if err != nil {
		return "", err
	}

	loginCode := random.Str(6)
	go accMan.notifyUser(email, loginCode)

	return loginCode, nil
}

func (accMan AccountManagement) notifyUser(email, loginCode string) {
	if err := accMan.sendNotification(email, loginCode); err != nil {
		logrus.Errorf("failed to send notification to user %s: %s", email, err.Error())
	}
}
