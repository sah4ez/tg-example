// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"

	"github.com/sah4ez/tg-example/pkg/interfaces"
)

type AdderAdd func(ctx context.Context, a int, b int) (c int, err error)

type MiddlewareAdder func(next interfaces.Adder) interfaces.Adder

type MiddlewareAdderAdd func(next AdderAdd) AdderAdd