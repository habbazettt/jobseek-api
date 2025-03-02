package config

import (
	"fmt"
	"log"
	"os"

	"github.com/habbazettt/jobseek-go/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	_ = godotenv.Load()

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("Error: Variabel DB_URL tidak ditemukan dalam .env")
	}

	// Koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Automigrate tabel berdasarkan model yang ada
	err = db.AutoMigrate(
		&models.User{},
		&models.Job{},
	)

	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}

	DB = db
	fmt.Println("Sukses terhubung ke database dan migrasi berhasil")
	return db
}
