package database

import (
	"fmt"
	"time"

	"backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(db config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		db.Host, db.User, db.Password, db.Name, db.Port,
	)

	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(db.MaxIdleConns)
	sqlDB.SetMaxOpenConns(db.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB, nil
}
