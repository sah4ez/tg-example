// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"

	"github.com/sah4ez/tg-example/pkg/interfaces"
)

type AdderSum func(ctx context.Context, aInt int, bInt int) (c int, err error)

type MiddlewareAdder func(next interfaces.Adder) interfaces.Adder

type MiddlewareAdderSum func(next AdderSum) AdderSum
