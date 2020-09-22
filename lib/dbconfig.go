package lib

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Gorm *gorm.DB

func init() {
	Gorm = gormDB()
}

func gormDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		"root", "root", "127.0.0.1", 3306, "test", true, "Local")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)

	_ = db.AutoMigrate(&New{})

	return db
}
