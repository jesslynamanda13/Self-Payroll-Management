package postgres

import (
	"fmt"
	"os"
	"self-payrol/model"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	log.Infof("initgorm")
	connection := os.Getenv("DATABASE_URL")
	fmt.Println(connection)

	// db, err := postgres.InitGorm()

	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		fmt.Printf("cant connect to database %s", err.Error())
	}

	err = db.AutoMigrate(&model.Position{}, &model.User{}, &model.Company{}, &model.Transaction{})

	if err != nil {
		fmt.Printf("migrate fail %s", err.Error())
	} else {
		fmt.Println("Database connection established successfully.")
	}

	return db

}
