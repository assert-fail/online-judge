package config

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name              string   `mapstructure:"name"`
		Port              int      `mapstructure:"port"`
		TrustedProxies    []string `mapstructure:"trusted_proxies"`
		Mode              string   `mapstructure:"mode"`
		JWTSecret         string   `mapstructure:"jwt_secret"`
		ServerExitTimeout int      `mapstructure:"server_exit_timeout"`
	} `mapstructure:"app"`
	Database struct {
		Postgres struct {
			Host         string `mapstructure:"host"`
			Port         int    `mapstructure:"port"`
			User         string `mapstructure:"user"`
			Password     string `mapstructure:"password"`
			Name         string `mapstructure:"name"`
			MaxIdleConns int    `mapstructure:"max_idle_conns"`
			MaxOpenConns int    `mapstructure:"max_open_conns"`
		} `mapstructure:"postgres"`
		Redis struct {
			Host         string `mapstructure:"host"`
			Password     string `mapstructure:"password"`
			DB           int    `mapstructure:"db"`
			MaxIdleConns int    `mapstructure:"max_idle_conns"`
			PoolSize     int    `mapstructure:"pool_size"`
			DialTimeout  int    `mapstructure:"dial_timeout"`
			ReadTimeout  int    `mapstructure:"read_timeout"`
			WriteTimeout int    `mapstructure:"write_timeout"`
		} `mapstructure:"redis"`
	} `mapstructure:"database"`
}

type Postgres struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Name         string `mapstructure:"name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type Redis struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	PoolSize     int    `mapstructure:"pool_size"`
	DialTimeout  int    `mapstructure:"dial_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

var AppConfig Config

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return nil, err
	}

	printConfig()
	return &AppConfig, nil
}

// 用于格式化配置信息的模板
const configTemplate = `✅ Configuration loaded:
😀 App:
	Name: {{.App.Name}}
	Port: {{.App.Port}}
	Trusted Proxies: {{range $index, $proxy := .App.TrustedProxies}}{{if $index}}, {{end}}{{$proxy}}{{end}}{{if not .App.TrustedProxies}}None{{end}}
	Mode: {{.App.Mode}}
	JWT Secret: {{.App.JWTSecret}}
	Server Exit Timeout: {{.App.ServerExitTimeout}}s
😀 Database:
	Postgres:
		Port: {{.Database.Postgres.Port}}
		User: {{.Database.Postgres.User}}
		Password: {{.Database.Postgres.Password}}
		Name: {{.Database.Postgres.Name}}
		Max Idle Conns: {{.Database.Postgres.MaxIdleConns}}
		Max Open Conns: {{.Database.Postgres.MaxOpenConns}}
	Redis:
		Host: {{.Database.Redis.Host}}
		Password: {{.Database.Redis.Password}}
		DB: {{.Database.Redis.DB}}
		Max Idle Conns: {{.Database.Redis.MaxIdleConns}}
		Pool Size: {{.Database.Redis.PoolSize}}
		Dial Timeout: {{.Database.Redis.DialTimeout}}s
		Read Timeout: {{.Database.Redis.ReadTimeout}}s
		Write Timeout: {{.Database.Redis.WriteTimeout}}s
`

// 使用模板输出配置信息到日志
func printConfig() {
	templ, err := template.New("config").Parse(configTemplate)
	if err != nil {
		log.Println("⚠️ Failed to parse config template: ", err)
		return
	}

	if err := templ.Execute(os.Stdout, AppConfig); err != nil {
		log.Println("⚠️ Failed to execute config template: ", err)
		return
	}
}
