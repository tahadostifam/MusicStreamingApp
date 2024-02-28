package database

import (
	"github.com/tahadostifam/MusicStreamingApp/api/models"
	postgres "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var (
	once     sync.Once
	instance = gorm.DB{}
)

// This function is used to migrate all the tables that is defined in models and if it fails it should panic the process because database tables are one
// of the most low level services that is necessary for other services in runtime.
func migrateTables(db *gorm.DB) {
	migrateErr := db.AutoMigrate(&models.User{})
	if migrateErr != nil {
		panic(migrateErr)
	}
}

func Migrate() error {
	err := instance.Migrator().CreateTable(models.User{})
	if err != nil {
		return err
	}

	return nil
}

func Connect(dsn string) *gorm.DB {
	once.Do(func() {
		config := &gorm.Config{}

		db, err := gorm.Open(postgres.Open(dsn), config)
		if err != nil {
			panic(err)
		}

		instance = *db

		// Migrating tables
		migrateTables(&instance)
	})

	return &instance
}

func CreateTestDatabase() *gorm.DB {
	config := &gorm.Config{}

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), config)
	if err != nil {
		panic(err)
	}

	// Migrating tables
	migrateTables(db)

	return db
}
