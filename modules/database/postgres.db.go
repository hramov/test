package database

import (
	"fmt"

	model "testing/modules/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct{}

var DB *gorm.DB = nil

func (o *Gorm) Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", "localhost", "admin", "admin", "testing", "5432", "disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}

func (o *Gorm) GetConnection() *gorm.DB {
	return DB
}

func (o *Gorm) Migrate() error {
	err := DB.AutoMigrate(&model.User{})
	return err
}
