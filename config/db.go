package config

import (
	"log"
	"mini-project/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Environment variabel belum diisi")
	}

	// dsn := fmt.Sprintf(`user=%s password=%s dbname=%s port=%s sslmode=disable client_encoding=UTF8`, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal terkoneksi di database", err)
	}

	err = database.AutoMigrate(&models.Bioskop{})
	if err != nil {
		log.Fatal("Gagal melakukan migration database :", err)
	}

	DB = database
	log.Println("Berhasil terkoneksi ke Database")
}
