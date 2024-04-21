package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
}

func MustLoad() *Config {
	path := getConfigPath()
	return MustLoadByPath(path)
}

func getConfigPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("")
	}
	return path
}

func MustLoadByPath(path string) *Config {
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Sprintf("the file with path %s does not exist", path))
	}
	var config *Config
	if err := cleanenv.ReadConfig(path, config); err != nil {
		panic("error parsing config file")
	}
	return config
}
