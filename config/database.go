package config

import (
	"ddd/domain/model"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func ConnectDB() *gorm.DB {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// test:test@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=true", user, pass, dbName)

	// コネクションプールの生成
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Enable Connect to DB: %v", err)
	} else {
		fmt.Println("Successfully Connected DB")
	}

	// テーブルの作成
	return db.AutoMigrate(&model.User{}, &model.Habit{})

}
