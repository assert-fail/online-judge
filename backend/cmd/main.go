package main

import (
	"backend/config"
	"backend/pkg/database"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("❌ Unable to read config file: ", err)
	}

	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatalln("❌ Unable to initialize database: ", err)
	}

	if err := db.AutoMigrate(); err != nil {
		log.Fatalln("❌ Unable to migrate database: ", err)
	}
}
