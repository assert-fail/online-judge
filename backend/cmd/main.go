package main

import (
	"backend/config"
	"backend/internal/middleware"
	"backend/internal/models/user"
	"backend/internal/pkg/database"
	"backend/internal/pkg/logger"
	"backend/internal/router"

	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("❌ Unable to read config file:", err)
	}

	// 初始化日志
	logger.Init(cfg.App.Mode)

	// 初始化数据库连接池
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("❌ Unable to initialize database")
	}
	if err := db.AutoMigrate(&user.User{}); err != nil {
		logger.Fatal().
			Err(err).
			Msg("❌ AutoMigrate failed")
	}

	// 设置 Gin 模式
	if cfg.App.Mode == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建 Gin 引擎
	r := gin.New()

	// 设置可信的IP
	if err := r.SetTrustedProxies(cfg.App.TrustedProxies); err != nil {
		logger.Fatal().
			Err(err).
			Msg("❌ Unable to set trusted proxies")
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

	// 启动服务器
	logger.Info().Msgf("Starting server on port %d", cfg.App.Port)

	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		logger.Fatal().
			Err(err).
			Msg("❌ Unable to start server")
	}
}
