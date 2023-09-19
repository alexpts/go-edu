package provider

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alexpts/edu-go/internal/model"
	"github.com/alexpts/edu-go/internal/repo"
)

func ProvideDbConnect(config *Config) (*sql.DB, error) {
	// @todo need env parser with envsubst for POSTGRES_ENV one variable
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow application_name=edu",
		config.GetString("POSTGRES_HOST"),
		config.GetString("POSTGRES_USER"),
		config.GetString("POSTGRES_PASSWORD"),
		config.GetString("POSTGRES_DBNAME"),
		config.GetInt("POSTGRES_PORT"),
	)

	return sql.Open("pgx", dsn)
}

func ProvideGormDb(connection *sql.DB, logger *zerolog.Logger) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: connection,
	}), &gorm.Config{
		//Logger: logger, // @todo need adapter https://github.com/moul/zapgorm2/blob/master/zapgorm2.go
	})

	_ = gormDB.AutoMigrate(&model.User{}, &model.Post{}, &model.Category{})
	configPool(gormDB)

	return gormDB, err
}

func configPool(gormDB *gorm.DB) {
	sqlDB, _ := gormDB.DB()
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}

func ProvideUserRepo(db *gorm.DB) *repo.User {
	r := &repo.User{}
	r.Db = db
	return r
}

func ProvidePostRepo(db *gorm.DB) *repo.Post {
	return &repo.Post{
		Repo: repo.Repo[model.Post]{
			Db: db,
		},
	}
}

func ProvideCategoryRepo(db *gorm.DB) *repo.Category {
	return &repo.Category{
		Repo: repo.Repo[model.Category]{
			Db: db,
		},
	}
}
