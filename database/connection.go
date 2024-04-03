package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
}

func ConnectToDB() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		log.Info().Err(err).Msg("Database connection failed")
		return nil, err
	}
	config := Config{
		host:     os.Getenv("HOST"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
		port:     os.Getenv("PORT"),
		sslmode:  os.Getenv("SSL_MODE"),
	}

	dsn := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s`,
		config.host, config.port, config.user, config.password, config.dbname, config.sslmode,
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info().Err(err).Msg("Database connection failed")
		return nil, err
	}
	log.Info().Msg("Database connection successful")

	return DB, err

}

func MigrateAll(db *gorm.DB) error {
	err := Aircrafts(db)
	if err != nil {
		return err
	}
	return nil
}
