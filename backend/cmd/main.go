package main

import (
	"backend/config"
	"backend/internal/router"
	"backend/pkg/database"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	r := router.SetupRouter(gin.DebugMode, cfg.App.TrustedProxies)
	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		log.Fatalln("❌ Unable to start server: ", err)
	}
}
