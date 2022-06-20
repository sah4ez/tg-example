package interfaces

import "context"

// @tg http-server log
// @tg http-prefix=api/v1
type Files interface {
	// @tg http-method=GET
	// @tg http-success=200
	// @tg http-path=/public/file/template
	GetTemplate(ctx context.Context) (data []byte, name string, err error)
}
