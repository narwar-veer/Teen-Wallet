package config

import (
    "flag"
    "log"
    "os"

    "github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
    Addr string `yaml:"address" env-required:"true"`
}

type Postgres struct {
    Host     string `yaml:"host" env-required:"true"`
    Port     int    `yaml:"port" env-required:"true"`
    User     string `yaml:"user" env-required:"true"`
    Password string `yaml:"password" env-required:"true"`
    DBName   string `yaml:"dbname" env-required:"true"`
    SSLMode  string `yaml:"sslmode" env-required:"true"`
}

type JWT struct {
    Secret string `yaml:"secret" env-required:"true"`
}

type Config struct {
    Env        string   `yaml:"env" env:"ENV" env-required:"true"`
    HTTPServer HTTPServer `yaml:"http_server"`
    Postgres   Postgres   `yaml:"postgres"`
    JWT        JWT        `yaml:"jwt"`
}

func MustLoad() *Config {
    var configPath string

    configPath = os.Getenv("CONFIG_PATH")
    if configPath == "" {
        flags := flag.String("config", "", "path to the configuration file")
        flag.Parse()
        configPath = *flags
        if configPath == "" {
            log.Fatal("config path is not set")
        }
    }

    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        log.Fatalf("config file does not exist: %s", configPath)
    }

    var cfg Config
    if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
        log.Fatalf("cannot read config: %v", err)
    }
    return &cfg
}
