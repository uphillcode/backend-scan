// package database

// import (
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// var DB *gorm.DB

// func InitDB() {
// 	dsn := "root:root@tcp(127.0.0.1:3306)/scanscore?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	DB = db

//		// Auto-migrate the schema
//		// db.AutoMigrate(&User{})
//	}
package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	// Auto-migrate the schema
	// db.AutoMigrate(&YourModel{})
}

func getDSN() string {
	user := os.Getenv("DB_TRAMITE_USER")
	pass := os.Getenv("DB_TRAMITE_PASS")
	host := os.Getenv("DB_TRAMITE_HOST")
	port := os.Getenv("DB_TRAMITE_PORT")
	dbname := os.Getenv("DB_TRAMITE_DB")
	return user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}
