package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"` // read from yaml file with key, or from environment variables
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config { // Must functions generally don't return error. they are do or die
	var configPath string
	configPath = os.Getenv("CONFIG_PATH") // Environment variable
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file") // command line arguments/flags
		flag.Parse()

		configPath = *flags // de-reference into config-path

		if configPath == "" {
			log.Fatal("config path is required. Set CONFIG_PATH environment variable or use -config flag")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) { // error of file
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("could not read the config file: %s", err.Error())
	}
	return &cfg
}
