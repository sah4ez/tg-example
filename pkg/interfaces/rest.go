package interfaces

import "context"

// @tg http-server log
// @tg http-prefix=api/v1
type Adder interface {
	// @tg http-method=POST
	// @tg http-success=200
	// @tg http-path=/sum
	Sum(ctx context.Context, aInt int, bInt int) (c int, err error)
}
