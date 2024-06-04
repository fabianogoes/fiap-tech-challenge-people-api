package repository

import (
	"context"
	"fmt"
	"github.com/fabianogoes/fiap-people/domain/entities"
	"github.com/fabianogoes/fiap-people/frameworks/repository/dbo"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(ctx context.Context, config *entities.Config) (*gorm.DB, error) {
	loc, _ := time.LoadLocation("UTC")

	var dsnTemplate string
	if config.Environment == "production" {
		dsnTemplate = "user=%s password=%s host=%s port=%s dbname=%s TimeZone=%s"
	} else {
		dsnTemplate = "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s"
	}

	dsn := fmt.Sprintf(dsnTemplate,
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		loc,
	)

	fmt.Printf("DB_HOST = %s\n", os.Getenv("DB_HOST"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	if err = db.AutoMigrate(
		&dbo.Customer{},
		&dbo.Attendant{},
	); err != nil {
		log.Fatal("AutoMigrate error", err)
		return nil, err
	}

	InitialDataAttendants(db)
	InitialDataCustomers(db)

	return db, nil
}
