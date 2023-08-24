package integrations

import (
	"encoding/json"
	"errors"
	"github.com/mahalahub/mahala/internal/redis"
	"github.com/mahalahub/mahala/user"
)

const (
	Accounts = "_accounts"
)

type UserRepository struct {
	redisClient *redis.Client
}

func NewUserRepository(redisClient *redis.Client) *UserRepository {
	return &UserRepository{
		redisClient: redisClient,
	}
}

func (repo *UserRepository) GetUserAccount(username string) (user.Account, error) {
	accountRaw, err := repo.redisClient.Get(username)
	if err != nil && !errors.Is(err, redis.ErrNotExists) {
		return user.Account{}, err
	}

	var acc user.Account
	if len(accountRaw) > 0 {
		if err = json.Unmarshal(accountRaw, &acc); err != nil {
			return user.Account{}, err
		}
	}

	return acc, nil
}

func (repo *UserRepository) AddUserAccount(account user.Account) (user.Account, error) {
	if err := repo.redisClient.Add(redis.Item{
		Key:   account.Username,
		Value: account,
	}); err != nil {
		return user.Account{}, err
	}

	return account, nil
}
