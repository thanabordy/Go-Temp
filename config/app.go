package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	godotenv.Load()
	app()
	database()
}

func env(key string, fallback interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func config(key string, value interface{}) {
	viper.Set(key, value)
}

func app() {
	config("APP_PORT", env("APP_PORT", 80))
	config("APP_GRPC_PORT", env("APP_GRPC_PORT", 8000))
	config("WALLET_GRPC_IP", env("WALLET_GRPC_IP", "0.0.0.0:9090"))

	config("SSL_PRIVATE_PATH", env("SSL_PRIVATE_PATH", "storage/cert/privatekey.key"))
	config("SSL_CERT_PATH", env("SSL_CERT_PATH", "storage/cert/certificate.crt"))
}
