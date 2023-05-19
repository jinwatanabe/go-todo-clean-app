package driver

import (
	"fmt"
	"go-todo-clean-app/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvideDatabaseConnection() *gorm.DB {
	dbConfig := config.GetConfig().Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/godo?charset=utf8&parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}