package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(dsn string) (*Storage, error) {

	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			},
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}
