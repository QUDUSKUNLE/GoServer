package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	domain "server/internal/core/domain"
)

type PostgresDatabase struct {
	db *gorm.DB
}

func NewPostgresDatabase(host, port, user, password, dbname string) *PostgresDatabase {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
  )
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	} else {
		fmt.Println("Successfully connected to the database.")
	}
	db.AutoMigrate(&domain.Address{}, &domain.User{})
	return &PostgresDatabase{
		db: db,
	}
}
