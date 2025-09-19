package main

import (
	"backend/config"
	"backend/internal/middleware"
	"backend/internal/router"
	"backend/pkg/database"
	"backend/pkg/logger"
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

	logger.Init(cfg.App.Mode)

	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatalln("❌ Unable to initialize database:", err)
	}
	if err := db.AutoMigrate(); err != nil {
		log.Fatalln("❌ Unable to migrate database:", err)
	}

	if cfg.App.Mode == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(
		middleware.RequestIDMiddleware(),
		middleware.ErrorHandler(),
		middleware.GinLogger(),
		middleware.GinRecovery(),
		gin.Recovery(),
	)
	router.SetupRouter(r)

	if err := r.SetTrustedProxies(cfg.App.TrustedProxies); err != nil {
		log.Fatalln("❌ Unable to set trusted proxies:", err)
	}

	logger.Info().Msgf("Starting server on port %d", cfg.App.Port)
	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		logger.Error().Err(err).Msg("❌ Unable to start server")
	}
}
