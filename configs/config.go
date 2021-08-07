package configs

import (
	"pembayaran/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/amartha?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed Connection Database")
	}
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&users.User{})
}
