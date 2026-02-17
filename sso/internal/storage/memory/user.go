package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/kuaukutsu/auth/sso/internal/domain/models"
	"github.com/kuaukutsu/auth/sso/internal/storage"
)

type tableUser struct {
	mu      sync.RWMutex
	user    map[string]*rowUser
	uiEmail map[string]*rowUser
}

type rowUser struct {
	uuid     string
	email    string
	passHash string
}

func NewUser() *tableUser {
	return &tableUser{
		user:    make(map[string]*rowUser),
		uiEmail: make(map[string]*rowUser),
	}
}

func (s *tableUser) SaveUser(
	ctx context.Context,
	uuid string,
	email string,
	passHash []byte,
) (string, error) {
	const op = "storage.memory.SaveUser"

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.user[uuid]; exists {
		return "", fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	if _, exists := s.uiEmail[email]; exists {
		return "", fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	row := rowUser{
		uuid:     uuid,
		email:    email,
		passHash: string(passHash),
	}

	s.user[uuid] = &row
	s.uiEmail[email] = &row

	return uuid, nil
}

func (s *tableUser) User(
	ctx context.Context,
	email string,
) (models.User, error) {
	const op = "storage.memory.User"

	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.uiEmail[email]
	if exists == false {
		return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return models.User{
		Uuid:     user.uuid,
		Email:    user.email,
		PassHash: []byte(user.passHash),
	}, nil
}
