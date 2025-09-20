package logger

import (
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	GlobalLogger zerolog.Logger
	logFile      *os.File
	onceLogger   sync.Once
)

func Init(env string) (*os.File, error) {
	onceLogger.Do(func() {
		zerolog.TimeFieldFormat = time.RFC3339Nano
		zerolog.TimestampFunc = func() time.Time {
			return time.Now().UTC()
		}

		if env == "development" {
			// 开发环境：彩色控制台输出
			output := zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: "2006-01-02 15:04:05.000",
				FormatLevel: func(i any) string {
					return "\x1b[36m" + i.(string) + "\x1b[0m" // 青色
				},
				FormatMessage: func(i any) string {
					return "\x1b[32m" + i.(string) + "\x1b[0m" // 绿色
				},
			}

			GlobalLogger = zerolog.New(output).
				With().
				Timestamp().
				Logger()

			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			// 生产环境：JSON 格式
			var err error
			logFileName := "log/" + time.Now().Format("2006-01-02") + ".log"
			logFile, err = os.OpenFile(
				logFileName,
				os.O_CREATE|os.O_APPEND|os.O_WRONLY,
				0644,
			)
			if err != nil {
				log.Fatal().Err(err).Msg("❌ Failed to open log file")
			}

			GlobalLogger = zerolog.New(logFile).
				With().
				Timestamp().
				Logger()

			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		// 替换标准库的 log
		log.Logger = GlobalLogger
	})

	return logFile, nil
}

// 便捷方法
func Debug() *zerolog.Event {
	return GlobalLogger.Debug()
}

func Info() *zerolog.Event {
	return GlobalLogger.Info()
}

func Warn() *zerolog.Event {
	return GlobalLogger.Warn()
}

func Error() *zerolog.Event {
	return GlobalLogger.Error()
}

func Fatal() *zerolog.Event {
	return GlobalLogger.Fatal()
}

// WithRequest 创建带有请求上下文的 Logger
func WithRequest(c *gin.Context) zerolog.Logger {
	if c == nil {
		return GlobalLogger
	}

	return GlobalLogger.With().
		Str("request_id", getRequestID(c)).
		Str("client_ip", c.ClientIP()).
		Logger()
}

// 从 Gin 上下文获取请求 ID
func getRequestID(c *gin.Context) string {
	if requestID, exists := c.Get("X-Request-ID"); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return "unknown"
}
