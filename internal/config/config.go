package config

import (
	"flag"
	"fmt"
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

	RepositoryMode string `yaml:"repository-mode" env-default:"in-memory"`

	Postgres PostgresCfg `yaml:"postgres" env-required:"false"`
}

type PostgresCfg struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Migrate  bool   `yaml:"migrate"`
	Sslmode  string `yaml:"sslmode" env-default:"disable"`
}

func LoadConfig() *Cfg {
	cfg := Cfg{}

	flag.Parse()
	if err := cleanenv.ReadConfig(configPath(), &cfg); err != nil {
		log.Fatalf("[ERROR] read config %s", err.Error())
	}
	if cfg.RepositoryMode == "postgres" {
		validatePostgresConfig(cfg.Postgres)
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

func validatePostgresConfig(cfg PostgresCfg) error {
	if cfg.Host == "" ||
		cfg.Port == "" ||
		cfg.Password == "" ||
		cfg.Name == "" ||
		cfg.User == "" {
		return fmt.Errorf("config fields are required when using 'postgres' repository mode %+v", cfg)
	}
	return nil
}
