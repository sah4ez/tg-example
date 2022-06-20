package files

import (
	"context"
)

type Service struct{}

func (s *Service) GetTemplate(ctx context.Context) (data []byte, name string, err error) {
	data = Static.MustBytes("test.txt")
	name = "test.txt"

	return
}

func New() *Service {
	return &Service{}
}
