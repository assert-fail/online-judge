package config

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name           string   `mapstructure:"name"`
		Port           int      `mapstructure:"port"`
		TrustedProxies []string `mapstructure:"trusted_proxies"`
	} `mapstructure:"app"`
	Database struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		Name         string `mapstructure:"name"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
	} `mapstructure:"database"`
}

var appConfig Config

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, err
	}

	printConfig()
	return &appConfig, nil
}

// 用于格式化配置信息的模板
const configTemplate = `✅ Configuration loaded:
😀 App:
	Name: {{.App.Name}}
	Port: {{.App.Port}}
	Trusted Proxies: {{range $index, $proxy := .App.TrustedProxies}}{{if $index}}, {{end}}{{$proxy}}{{end}}{{if not .App.TrustedProxies}}None{{end}}
😀 Database:
	Port: {{.Database.Port}}
	User: {{.Database.User}}
	Password: {{.Database.Password}}
	Name: {{.Database.Name}}
	Max Idle Conns: {{.Database.MaxIdleConns}}
	Max Open Conns: {{.Database.MaxOpenConns}}
`

// 使用模板输出配置信息到日志
func printConfig() {
	templ, err := template.New("config").Parse(configTemplate)
	if err != nil {
		log.Println("⚠️ Failed to parse config template: ", err)
		return
	}

	if err := templ.Execute(os.Stdout, appConfig); err != nil {
		log.Println("⚠️ Failed to execute config template: ", err)
		return
	}
}
