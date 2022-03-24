package storage

import (
	"context"

	"github.com/sah4ez/tg-example/pkg/storage/types"
)

type User interface {
	GetUserByID(ctx context.Context, id int) (user types.User, err error)
}
