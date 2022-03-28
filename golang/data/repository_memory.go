package data

import (
	"fmt"
	"log"

	"github.com/guionardo/auth-service/golang/domain"
)

type RepositoryMemory struct {
	users  map[string]*domain.UserData
	tokens map[string]*domain.UserData
}

func (r *RepositoryMemory) Setup(args interface{}) {
	r.users = make(map[string]*domain.UserData)
}
func (r *RepositoryMemory) SetUser(user *domain.UserData) error {
	r.users[user.UserID] = user
	return nil
}
func (r *RepositoryMemory) GetUser(userId string) (*domain.UserData, error) {
	if user, ok := r.users[userId]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user not found '%s'", userId)
}

func (r *RepositoryMemory) DeleteUser(userId string) error {
	delete(r.users, userId)
	var tokens = make([]string, 0)
	for key, value := range r.tokens {
		if value.UserID == userId {
			tokens = append(tokens, key)
		}
	}
	for _, token := range tokens {
		delete(r.tokens, token)
	}
	return nil
}

func (r *RepositoryMemory) ValidatePassword(userData domain.UserData, password string) error {
	passwordHash := hasher.Hash(password)
	if userData.PasswordHash == passwordHash {
		return nil
	}
	return fmt.Errorf("INVALID PASSWORD")
}

func (r *RepositoryMemory) Show() {
	log.Println("Repository: IN MEMORY")
}
