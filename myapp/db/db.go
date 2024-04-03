package db

import (
	"myapp/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDatabase(dsn string) {
	var err error
	// dsn := "postgresql://user:password@localhost/database_name?sslmode=disable" // Update with your database credentials

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrate the User model
	Db.AutoMigrate(&entities.Board{}, &entities.User{})
}
