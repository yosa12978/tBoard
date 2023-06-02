package db

import (
	"fmt"
	"os"
	"sync"

	"github.com/yosa12978/tBoard/internal/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		connstr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		db, err = gorm.Open(postgres.Open(connstr), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		migrate()
	})
	return db
}

func migrate() {
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Account{})
}
