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
    Env        string     `yaml:"env" env:"ENV" env-required:"true"`
    HTTPServer HTTPServer `yaml:"http_server"`
    Postgres   Postgres   `yaml:"postgres"`
    JWT        JWT        `yaml:"jwt"`
}

// MustLoad loads configuration from YAML file. Order of precedence:
// 1. CONFIG_PATH environment variable
// 2. -config CLI flag (defaults to ./configs/local.yaml)
// 3. Built‑in default ./configs/local.yaml
// It terminates the program on any error.
func MustLoad() *Config {
    var configPath string

    // 1) Highest priority: explicit env var
    configPath = os.Getenv("CONFIG_PATH")

    // 2) CLI flag (overrides default)
    if configPath == "" {
        flag.StringVar(&configPath, "config", "./configs/local.yaml", "path to the configuration file")
        flag.Parse()
    }

    // 3) If still empty (possible when flag.Parse not called), use default
    if configPath == "" {
        configPath = "./configs/local.yaml"
    }

    // Ensure the file exists
    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        log.Fatalf("config file does not exist: %s", configPath)
    }

    var cfg Config
    if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
        log.Fatalf("cannot read config: %v", err)
    }

    return &cfg
}
