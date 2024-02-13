package database

import (
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	once     sync.Once
	instance = gorm.DB{}
)

func Connect() *gorm.DB {
	once.Do(func() {
		dsn := ""
		config := &gorm.Config{}

		db, err := gorm.Open(postgres.Open(dsn), config)
		if err != nil {
			panic(err)
		}

		instance = *db
	})

	return &instance
}
