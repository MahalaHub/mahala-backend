package user_test

import (
	"github.com/mahalahub/mahala/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountManagement_GenerateLoginCode_UserNotExists(t *testing.T) {
	repo := &repositoryMock{}
	accMan := user.NewAccountManagement(repo, sendNotificationMock)

	code, err := accMan.GenerateLoginCode("user-1", "email-1")
	assert.NoError(t, err)

	assert.Equal(t, 6, len(code))
	assert.Equal(t, 1, len(repo.accounts))
}

func TestAccountManagement_GenerateLoginCode_UserAlreadyExists(t *testing.T) {
	repo := &repositoryMock{
		accounts: []user.Account{
			{
				ID:       1,
				Username: "user-1",
				Email:    "email-1",
			},
		},
	}
	accMan := user.NewAccountManagement(repo, sendNotificationMock)

	code, err := accMan.GenerateLoginCode("user-1", "email-1")
	assert.Error(t, err)

	assert.Equal(t, 0, len(code))
}

type repositoryMock struct {
	accounts []user.Account
}

func (repo *repositoryMock) GetUserAccount(username string) (user.Account, error) {
	for _, acc := range repo.accounts {
		if acc.Username == username {
			return acc, nil
		}
	}
	return user.Account{}, nil
}

func (repo *repositoryMock) AddUserAccount(account user.Account) (user.Account, error) {
	acc := user.Account{
		ID:       len(repo.accounts) + 1,
		Username: account.Username,
		Email:    account.Email,
	}
	repo.accounts = append(repo.accounts, acc)
	return acc, nil
}

func sendNotificationMock(to, message string) error {
	return nil
}
