package config

import (
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/grpc"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/http"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/kafka"
	"github.com/D1sordxr/fin-eventor-lite/internal/infrastructure/config/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Storage       postgres.Config `yaml:"storage"`
	MessageBroker kafka.Config    `yaml:"message_broker"`
	HttpServer    http.Config     `yaml:"http"`
	GrpcServer    grpc.Config     `yaml:"grpc"`
}

const BasicConfigPath = "./configs/app/local.yaml"

func NewConfig() *Config {
	var cfg Config

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = BasicConfigPath
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
