package utils

import (
	"ecommerce/configs"
	"ecommerce/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitMigrate(db)

	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.ShoppingCart{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Payment{})
	db.AutoMigrate(&entities.Order{})
	db.AutoMigrate(&entities.Category{})
}