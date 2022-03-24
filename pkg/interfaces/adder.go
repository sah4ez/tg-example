package interfaces

import "context"

// @tg jsonRPC-server log
// @tg http-prefix=api/v1
type User interface {
	GetUserNameByID(ctx context.Context, id int) (name string, err error)
}
