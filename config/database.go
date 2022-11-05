package config

import (
	"ddd/domain/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB
// var err error

func ConnectDB() *gorm.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("Unable load location")
	}

	dsn := mysql.Config{
		User:      user,
		Passwd:    password,
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    dbName,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
		ParseTime: true,
	}

	// コネクションプールの生成
	db, err := gorm.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatalf("Unable connect to DB %s", err)
	} else {
		fmt.Println("Successfully Connected DB")
	}

	// テーブルの作成
	return db.AutoMigrate(&model.User{}, &model.Habit{})

}
