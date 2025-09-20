package main

import (
	"backend/config"
	"backend/internal/middleware"
	"backend/internal/models/user"
	"backend/internal/pkg/database"
	"backend/internal/pkg/logger"
	"backend/internal/router"
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cfg         *config.Config
	db          *gorm.DB
	redisClient *redis.Client
	logFile     *os.File
	r           *gin.Engine
)

func init() {
	log.SetOutput(os.Stdout)

	var err error
	// 加载配置
	cfg, err = config.LoadConfig()
	if err != nil {
		log.Fatalln("❌ Unable to read config file:", err)
	}

	// 初始化日志
	if logFile, err = logger.Init(cfg.App.Mode); err != nil {
		log.Fatalln("❌ Unable to initialize logger:", err)
	}

	// 初始化数据库连接池
	db, err = database.InitDB(cfg.Database.Postgres)
	if err != nil {
		logger.Fatal().Err(err).Msg("❌ Unable to initialize database")
	}
	if err = db.AutoMigrate(&user.User{}); err != nil {
		logger.Fatal().Err(err).Msg("❌ AutoMigrate failed")
	}

	// 初始化 Redis 连接池
	if redisClient, err = database.InitRedis(cfg.Database.Redis); err != nil {
		logger.Fatal().Err(err).Msg("❌ Unable to initialize Redis")
	}

	// 设置 Gin 模式
	if cfg.App.Mode == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建 Gin 引擎
	r = gin.New()

	// 设置可信的IP
	if err := r.SetTrustedProxies(cfg.App.TrustedProxies); err != nil {
		logger.Fatal().Err(err).Msg("❌ Unable to set trusted proxies")
	}

	// 添加中间件
	r.Use(
		middleware.RequestIDMiddleware(), // id -> next
		middleware.GinRecovery(),         // defer recover() -> next
		// gin.Recovery(),
		middleware.GinLogger(),    // time -> next -> log
		middleware.ErrorHandler(), // next -> err -> handle
	)

	// 创建Handler
	user := user.New(user.NewUserService(user.NewUserRepository(db)))

	// 设置 Gin 路由
	router.SetupRouter(r, user)
}

func close() {
	logger.Info().Msg("Closing resources...")

	if err := redisClient.Close(); err != nil {
		logger.Error().Err(err).Msg("❌ Error closing Redis")
	}

	if sqlDB, err := db.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			logger.Error().Err(err).Msg("❌ Error closing database")
		}
	}

	if logFile != nil {
		if err := logFile.Close(); err != nil {
			logger.Error().Err(err).Msg("❌ Error closing log file")
		}
	}

	log.Println("All resources closed.")
}

func main() {
	defer close()

	// 启动服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: r,
	}
	logger.Info().Msgf("Starting server on port %d", cfg.App.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("❌ Server failed to start")
		}
	}()

	// 等待中断信号以优雅关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.App.ServerExitTimeout)*time.Second,
	)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Warn().Err(err).Msg("⚠️ Server forced to shutdown")
	}

	logger.Info().Msg("Server exited successfully")
}
