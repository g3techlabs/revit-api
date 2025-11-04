package db

import (
	"fmt"

	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect(logger utils.ILogger) {
	logger.Info("Connecting to the database...")

	var err error

	dbHost := config.Get("DATABASE_HOST")
	dbUser := config.Get("DATABASE_USER")
	dbPassword := config.Get("DATABASE_PASSWORD")
	dbSchema := config.Get("DATABASE_SCHEMA")
	dbPort := config.GetIntVariable("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", dbHost, dbUser, dbPassword, dbSchema, dbPort)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Fatalf("Database connection gone wrong: %v", err)
	}

	logger.Info("Successfully connected to the database")
}
