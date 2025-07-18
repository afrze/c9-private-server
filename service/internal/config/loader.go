package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func Load() Config {
	v := viper.New()

	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.BindEnv("MONGODB_URI")
	v.BindEnv("MSSQL_CONN")
	v.BindEnv("JWT_SECRET")
	v.BindEnv("GIN_MODE")

	v.SetDefault("LISTEN_ADDR", ":9009")
	v.SetDefault("JWT_EXPIRY_MIN", 60)
	v.SetDefault("RATE_LIMIT", 60)
	v.SetDefault("GIN_MODE", "debug")

	_ = v.ReadInConfig()

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("config: %v", err)
	}
	if c.MongoURI == "" || c.MSSQLConn == "" || c.JWTSecret == "" {
		log.Fatalf("config: missing MONGODB_URI / MSSQL_CONN / JWT_SECRET")
	}
	return c
}
