package interfaces

import "context"

// @tg jsonRPC-server log
// @tg http-prefix=api/v1
type Adder interface {
	Add(ctx context.Context, a, b int) (c int, err error)
}
