package provider

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideDbConnect(config *Config) (*sql.DB, error) {
	// @todo need env parser with envsubst for POSTGRES_ENV one variable
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow",
		config.GetString("POSTGRES_HOST"),
		config.GetString("POSTGRES_USER"),
		config.GetString("POSTGRES_PASSWORD"),
		config.GetString("POSTGRES_DBNAME"),
		config.GetInt("POSTGRES_PORT"),
	)

	return sql.Open("pgx", dsn)
}

func ProvideGormDb(connection *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: connection,
	}), &gorm.Config{})

	//autoMigrate(gormDB)
	return gormDB, err
}
