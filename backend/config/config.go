package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// loadEnv memuat konfigurasi lingkungan dari file .env
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Falling back to system environment variables.")
	}
}

// getEnv mengambil variabel lingkungan dengan fallback ke nilai default
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		if fallback == "" {
			log.Fatalf("Environment variable %s is required but not set", key)
		}
		return fallback
	}
	return value
}

// SetupDB menghubungkan aplikasi dengan database MySQL menggunakan Gorm
func SetupDB() {
	loadEnv()

	// Ambil konfigurasi database dari .env file
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Format Data Source Name (DSN) untuk MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)

	// Membuka koneksi ke database MySQL menggunakan Gorm
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connected!")
}
