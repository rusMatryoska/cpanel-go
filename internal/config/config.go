package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8000"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type Storage struct {
	Address  string `yaml:"address" env-required:"true"`
	DBName   string `yaml:"db_name" env-required:"true"`
	Schema   string `yaml:"schema" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true" env:"DB_PASSWORD"`
}

func MustLoad() *Config {
	cfg := &Config{}
	os.Setenv("CONFIG_PATH", "/home/victoria/Desktop/cpanel-go/config/local.yaml") //TODO: delete
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	// check readable of file
	if yamlFile, err := os.ReadFile(configPath); err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	} else {
		// if file is readable, do unmarshall
		if err := yaml.Unmarshal(yamlFile, cfg); err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}

	return cfg
}
