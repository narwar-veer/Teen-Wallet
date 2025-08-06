package models

import "time"

type Wallet struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"uniqueIndex"`
    Balance   int64     `gorm:"default:0"` 
    Limit     int64     `gorm:"default:0"` 
    CreatedAt time.Time
    UpdatedAt time.Time
}
