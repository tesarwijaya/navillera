package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config add your desired environment variable here
type Config struct {
	AppName  string `envconfig:"app_name"`
	AppQuote string `envconfig:"app_quote"`
	GrpcHost string `envconfig:"grpc_host"`
	GrpcPort string `envconfig:"grpc_port"`
	LogLevel string `envconfig:"log_level"`
	Version  string `envconfig:"version"`
}

// C ...
var C Config

func init() {
	err := envconfig.Process("navillera", &C)
	if err != nil {
		log.Fatal(err.Error())
	}
}
