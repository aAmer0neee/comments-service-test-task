package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	config_path_flag = "config-path"
	config_path_env  = "CONFIG_PATH"
)

var (
	path = flag.String(config_path_flag, "", "path to configure file")
)

type Cfg struct {
	Server struct {
		Host string `yaml:"host" env-required:"true"`
		Port string `yaml:"port" env-required:"true"`
		Env  string `yaml:"env" env-default:"local"`
	} `yaml:"server" env-required:"true"`

	RepositoryMode string `yaml:"repository-mode" env-default:"postgres"`

	Postgres struct {
		Port     string `yaml:"port" env-required:"true"`
		Host     string `yaml:"host" env-required:"true"`
		Name     string `yaml:"name" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		User     string `yaml:"user" env-required:"true"`
		Migrate  bool   `yaml:"migrate" env-default:"false"`
		Sslmode  string `yaml:"sslmode" env-default:"disable"`
	} `yaml:"postgres" env-required:"true"`
}

func LoadConfig() *Cfg {
	cfg := Cfg{}

	flag.Parse()
	if err := cleanenv.ReadConfig(configPath(), &cfg); err != nil {
		log.Fatalf("[ERROR] read config %s", err.Error())
	}

	return &cfg
}

func configPath() string {

	if *path == "" {
		*path = os.Getenv(config_path_env)
	}
	if _, err := os.Stat(*path); err == os.ErrNotExist {
		log.Fatalf("[ERROR] no such file %s", *path)
	}

	return *path
}
