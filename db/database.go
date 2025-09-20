package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/g3techlabs/revit-api/db/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Setup() {

	dbHost := os.Getenv("DATABASE_HOST")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbSchema := os.Getenv("DATABASE_SCHEMA")
	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	if err != nil {
		fmt.Printf("%s", dbUser)
		panic("Error converting dbPort to integer type" + err.Error())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", dbHost, dbUser, dbPassword, dbSchema, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("Database connection gone wrong" + err.Error())
	}

	db.AutoMigrate(&entities.User{})
}
