package interfaces

import "context"

// @tg http-server log
// @tg http-prefix=api/v1
type Adder interface {
	// @tg http-method=GET
	// @tg http-success=200
	// @tg http-path=/sum
	// @tg http-args=aInt|a,bInt|b
	Sum(ctx context.Context, aInt int, bInt int) (c int, err error)
}
