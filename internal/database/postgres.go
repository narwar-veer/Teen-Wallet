package database

import (
    "fmt"
    "log"

    "github.com/yourusername/teen-wallet-api/internal/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
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

// AutoMigrate auto‑creates tables
func AutoMigrate(db *gorm.DB) {
    if err := db.AutoMigrate(&User{}, &Wallet{}, &Transaction{}); err != nil {
        log.Fatalf("auto‑migration failed: %v", err)
    }
}

================ internal/models/user.go ===================
package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey"`
    Name         string    `gorm:"size:100;not null"`
    Email        string    `gorm:"size:100;uniqueIndex"`
    Phone        string    `gorm:"size:15;uniqueIndex"`
    PasswordHash string    `gorm:"size:255;not null"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    Wallet       Wallet
}