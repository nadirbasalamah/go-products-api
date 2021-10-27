package database

import (
	"fmt"

	"github.com/nadirbasalamah/go-products-api/config"
	"github.com/nadirbasalamah/go-products-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	var err error

	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected!")

	DB.AutoMigrate(&model.Product{})
}
