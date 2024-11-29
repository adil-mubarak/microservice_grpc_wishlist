package db

import (
	"log"
	"microservice_grpc_wishlist/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "root:kl18jda183079@tcp(127.0.0.1:3306)/wishlist_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect with database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&models.Wishlist{})
	if err != nil {
		log.Fatalf("failed to migrate the table: %v", err)
		return nil, err
	}

	return db, nil
}
