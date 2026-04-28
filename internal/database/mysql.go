package database

import (
	"fmt"
	"log"

	"github.com/vibecart/vibecart-server/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=skip-verify",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect MySQL: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping MySQL: ", err)
	}

	log.Println("MySQL connected and ping successfully")
	return db
}