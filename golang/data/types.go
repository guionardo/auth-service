package data

import (
	"time"

	"github.com/guionardo/auth-service/golang/domain"
)

type (
	Cache interface {
		Setup(args interface{}) Cache
		Set(key string, value interface{}, timeToLive time.Duration) error
		Get(key string) (interface{}, error)
	}

	Repository interface {
		Setup(args interface{})
		SetUser(user *domain.UserData) error
		GetUser(userId string) (*domain.UserData, error)
		DeleteUser(userId string) error
		ValidatePassword(userData domain.UserData, password string) error
		Show()
	}
)
