package database

import (
    "fmt"
    "log"

    "github.com/narwar-veer/teen-wallet-api/internal/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/narwar-veer/teen-wallet-api/internal/models"
)

func MustConnect(cfg *config.Config) *gorm.DB {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.Postgres.Host,
        cfg.Postgres.Port,
        cfg.Postgres.User,
        cfg.Postgres.Password,
        cfg.Postgres.DBName,
        cfg.Postgres.SSLMode,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    return db
}

func AutoMigrate(db *gorm.DB) {
    if err := db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{}); err != nil {
        log.Fatalf("auto‑migration failed: %v", err)
    }
}
