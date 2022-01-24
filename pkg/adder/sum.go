package adder

import "context"

type Service struct {
}

func (s *Service) Add(_ context.Context, a, b int) (c int, err error) {
	return a + b, nil
}

func New() *Service {
	return &Service{}
}
