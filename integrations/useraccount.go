package integrations

import "github.com/mahalahub/mahala/user"

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (repo UserRepository) GetUserAccount(username string) (user.Account, error) {
	return user.Account{}, nil
}

func (repo UserRepository) AddUserAccount(account user.Account) (user.Account, error) {
	return user.Account{}, nil
}
