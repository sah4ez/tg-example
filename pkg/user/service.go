package user

import (
	"context"

	"github.com/sah4ez/tg-example/pkg/errors"
	"github.com/sah4ez/tg-example/pkg/storage"
)

type Service struct {
	userStore storage.User
}

func (s *Service) GetUserNameByID(ctx context.Context, id int) (name string, err error) {

	u, err := s.userStore.GetUserByID(ctx, id)
	if err != nil {
		err = errors.ErrUserNotFound.SetCause(err.Error())
		return
	}

	name = u.Name
	return
}

func New(userStore storage.User) *Service {
	return &Service{
		userStore: userStore,
	}
}
