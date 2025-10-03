package db

import (
	"fmt"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/users/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect() {
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
		panic("Database connection gone wrong" + err.Error())
	}

	fmt.Println("Successfully connected to the database")

	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error during migrations:" + err.Error())
	}
	fmt.Println("Database successfully migrated")
}
