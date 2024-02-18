package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Env is 環境変数
type Env struct {
	AppEnv     string `envconfig:"APP_ENV" default:"dev"`
	DBHost     string `envconfig:"DB_HOST" default:"rdb"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBName     string `envconfig:"DB_NAME" default:"cosmo"`
	DBUserName string `envconfig:"DB_USERNAME" default:"postgres"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"cosmo"`
	SSLMode    string `envconfig:"SSL_MODE" default:"disable"`
	LogFormat  string `envconfig:"LOG_FORMAT" default:"json"`
	LogLevel   string `envconfig:"LOG_LEVEL" default:"debug"`
}

var env Env

func init() {
	if err := envconfig.Process("", &env); err != nil {
		panic(fmt.Errorf("failed to get environment variables: %w", err))
	}
}

func GetEnv() *Env {
	return &env
}

// IsTest : Check if the current environment is test
func IsTest() bool {
	return env.AppEnv == "test"
}

// IsDev : Check if the current environment is development
func IsDev() bool {
	return env.AppEnv == "dev"
}

// IsStg : Check if the current environment is staging
func IsStg() bool {
	return env.AppEnv == "stg"
}

// IsPrd : Check if the current environment is production
func IsPrd() bool {
	return env.AppEnv == "prd"
}

// func GetOAuthGoogleRedirectURL() string {
// 	switch {
// 	case IsLocal():
// 		return "http://localhost:1323/oauth/google/callback"
// 	case IsTest():
// 		return "" // TODO
// 	case IsDev():
// 		return "" // TODO
// 	case IsStg():
// 		return "" // TODO
// 	case IsPrd():
// 		return "" // TODO
// 	}
// 	return ""
// }
