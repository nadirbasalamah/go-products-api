package database

import (
	"errors"
	"fmt"

	"github.com/nadirbasalamah/go-products-api/config"
	"github.com/nadirbasalamah/go-products-api/model"
	"github.com/nadirbasalamah/go-products-api/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dbName string) {
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), dbName)
	var err error

	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected!")

	DB.AutoMigrate(&model.Product{})

	DB.AutoMigrate(&model.User{})
}

func SeedProduct() (model.Product, error) {
	product, err := utils.CreateFakerData[model.Product]()
	if err != nil {
		return model.Product{}, nil
	}

	DB.Create(&product)
	fmt.Println("Product seeded to the database")

	return product, nil
}

func SeedUser() (model.User, error) {
	user, err := utils.CreateFakerData[model.User]()
	if err != nil {
		return model.User{}, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return model.User{}, err
	}

	var inputUser model.User = model.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: string(password),
	}

	DB.Create(&inputUser)
	fmt.Println("User seeded to the database")

	return user, nil
}

func CleanSeeders() {
	productResult := DB.Exec("TRUNCATE products")
	userResult := DB.Exec("TRUNCATE users")

	var isFailed bool = productResult.Error != nil || userResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	fmt.Println("Seeders are cleaned up successfully")
}
