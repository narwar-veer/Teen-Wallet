package models

import "time"

type TransactionType string

const (
    Deposit    TransactionType = "deposit"
    Withdraw   TransactionType = "withdraw"
    TransferIn TransactionType = "transfer_in"
    TransferOut TransactionType = "transfer_out"
)

type Transaction struct {
    ID          uint            `gorm:"primaryKey"`
    WalletID    uint            `gorm:"index"`
    Amount      int64           // positive paise
    Type        TransactionType `gorm:"size:20"`
    Description string          `gorm:"size:255"`
    CreatedAt   time.Time
}