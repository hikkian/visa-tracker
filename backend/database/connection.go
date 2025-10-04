package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"visa-tracker/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty",
		host, user, password, dbname, port,
	)

	var database *gorm.DB

	for {
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Не удалось подключиться к БД:", err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}

	}
	database.AutoMigrate(&models.Migrant{})

	DB = database
	fmt.Println("PostgreSQL подключен успешно.")
}
