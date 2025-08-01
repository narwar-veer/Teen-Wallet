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