package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	APP_NAME          string = "APP_NAME"
	APP_PORT          string = "APP_PORT"
	DATABASE_HOST     string = "DATABASE_HOST"
	DATABASE_PORT     string = "DATABASE_PORT"
	DATABASE_NAME     string = "DATABASE_NAME"
	DATABASE_USERNAME string = "DATABASE_USERNAME"
	DATABASE_PASSWORD string = "DATABASE_PASSWORD"
)

const (
	ENV_LOCAL string = "local"
)

type Config struct {
	App AppConfig
	Db  DatabaseConfig
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	v := viper.New()
	bindAllEnv(v)
	v.AutomaticEnv()

	var app AppConfig
	if err := v.Unmarshal(&app); err != nil {
		panic(err)
	}
	var db DatabaseConfig
	if err := v.Unmarshal(&db); err != nil {
		panic(err)
	}
	return &Config{
		App: app,
		Db:  db,
	}
}

func bindAllEnv(v *viper.Viper) {
	envs := os.Environ()
	for _, env := range envs {
		key := strings.Split(env, "=")[0]
		v.BindEnv(key)
	}
}

type AppConfig struct {
	Name string `mapstructure:"APP_NAME"`
	Port int    `mapstructure:"APP_PORT"`
	Env  string `mapstructure:"APP_ENV"`
}

func (a *AppConfig) IsLocal() bool {
	return strings.Compare(ENV_LOCAL, a.Env) == 0
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     int    `mapstructure:"DATABASE_PORT"`
	Name     string `mapstructure:"DATABASE_NAME"`
	Username string `mapstructure:"DATABASE_USERNAME"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
}
