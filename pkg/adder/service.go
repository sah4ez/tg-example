package adder

import (
	"context"
)

type Service struct{}

func (s *Service) Sum(ctx context.Context, aInt int, bInt int) (c int, err error) {

	return aInt + bInt, nil
}

func New() *Service {
	return &Service{}
}
