package storage

import (
	"github.com/sah4ez/tg-example/pkg/storage/types"
)

func (s *Storage) Migrate() (err error) {

	err = s.db.AutoMigrate(tables...)
	if err != nil {
		return
	}

	return
}

var tables = []interface{}{
	&types.User{},
}
