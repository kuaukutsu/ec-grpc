package memory

import (
	"context"
	"fmt"

	"github.com/kuaukutsu/auth/sso/internal/domain/models"
	"github.com/kuaukutsu/auth/sso/internal/storage"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	app     map[int]AppRow
	user    map[string]*UserRow
	uiEmail map[string]*UserRow
}

type AppRow struct {
	id     int
	name   string
	secret string
}

type UserRow struct {
	uuid     string
	email    string
	passHash string
}

type UserIndexEmail map[string]UserRow

func New() *Storage {
	return &Storage{
		app:     make(map[int]AppRow),
		user:    make(map[string]*UserRow),
		uiEmail: make(map[string]*UserRow),
	}
}

func (s *Storage) SaveUser(
	ctx context.Context,
	uuid string,
	email string,
	passHash []byte,
) (string, error) {
	const op = "storage.memory.SaveUser"

	if _, exists := s.user[uuid]; exists {
		return "", fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	if _, exists := s.uiEmail[email]; exists {
		return "", fmt.Errorf("%s: %w", op, storage.ErrUserExists)
	}

	row := UserRow{
		uuid:     uuid,
		email:    email,
		passHash: string(passHash),
	}

	s.user[uuid] = &row
	s.uiEmail[email] = &row

	return uuid, nil
}

func (s *Storage) User(
	ctx context.Context,
	email string,
) (models.User, error) {
	const op = "storage.memory.User"

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

func (s *Storage) App(
	ctx context.Context,
	id int,
) (models.App, error) {
	const op = "storage.memory.App"

	if id == 1 {
		return models.App{
			ID:     1,
			Name:   "test",
			Secret: "718e4894-a518-4802-9205-4838c7ddbd42",
		}, nil
	}

	app, exists := s.app[id]
	if exists == false {
		return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	}

	return models.App{
		ID:     app.id,
		Name:   app.name,
		Secret: app.secret,
	}, nil
}
