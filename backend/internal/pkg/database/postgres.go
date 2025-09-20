package database

import (
	"fmt"
	"sync"
	"time"

	"backend/config"
	"backend/internal/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance   *gorm.DB
	oncePostgres sync.Once
)

func InitDB(db config.Postgres) (*gorm.DB, error) {
	oncePostgres.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			db.Host, db.User, db.Password, db.Name, db.Port,
		)

		var err error
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Fatal().Err(err).Msg("❌ Failed to connect to database")
		}

		sqlDB, err := dbInstance.DB()
		if err != nil {
			logger.Fatal().Err(err).Msg("❌ Failed to get database instance")
		}

		sqlDB.SetMaxIdleConns(db.MaxIdleConns)
		sqlDB.SetMaxOpenConns(db.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Hour)

		logger.Info().Msg("PostgreSQL connected")
	})

	return dbInstance, nil
}
