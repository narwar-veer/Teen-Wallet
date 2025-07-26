package models

import "time"

type Wallet struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"uniqueIndex"`
    Balance   int64     `gorm:"default:0"` // paise
    Limit     int64     `gorm:"default:0"` // daily spending limit
    CreatedAt time.Time
    UpdatedAt time.Time
}
