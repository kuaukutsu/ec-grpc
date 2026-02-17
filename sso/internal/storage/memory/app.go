package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/kuaukutsu/auth/sso/internal/domain/models"
	"github.com/kuaukutsu/auth/sso/internal/storage"
)

type tableApp struct {
	mu  sync.RWMutex
	app map[int]rowApp
}

type rowApp struct {
	id     int
	name   string
	secret string
}

func NewApp() *tableApp {
	return &tableApp{
		app: make(map[int]rowApp),
	}
}

func (s *tableApp) App(
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

	s.mu.RLock()
	defer s.mu.RUnlock()

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
