package storage

import (
	"context"

	"github.com/sah4ez/tg-example/pkg/storage/types"
)

func (s *Storage) GetUserByID(ctx context.Context, id int) (user types.User, err error) {

	err = s.db.WithContext(ctx).Model(&user).Where("id = ?", id).First(&user).Error

	return
}
