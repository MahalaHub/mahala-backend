package integrations

import "github.com/mahalahub/mahala/user"

type UserRepository struct {
	accounts []user.Account
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) GetUserAccount(username string) (user.Account, error) {
	for _, acc := range repo.accounts {
		if acc.Username == username {
			return acc, nil
		}
	}
	return user.Account{}, nil
}

func (repo *UserRepository) AddUserAccount(account user.Account) (user.Account, error) {
	account.ID = len(repo.accounts) + 1
	repo.accounts = append(repo.accounts, account)
	return account, nil
}
