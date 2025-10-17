package db

import (
	"fmt"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() {
	utils.Log.Info("Connecting to the database...")

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
		utils.Log.Fatalf("Database connection gone wrong: %v", err)
	}

	utils.Log.Info("Successfully connected to the database")
}
