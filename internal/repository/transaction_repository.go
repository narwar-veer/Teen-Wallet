package repository

import (
    "github.com/narwar-veer/teen-wallet-api/internal/models"
    "gorm.io/gorm"
)

type TransactionRepository interface {
    Create(tx *models.Transaction) error
    GetByWallet(walletID uint, limit int) ([]models.Transaction, error)
}

type transactionRepo struct { db *gorm.DB }

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
    return &transactionRepo{db: db}
}

func (r *transactionRepo) Create(txn *models.Transaction) error {
    return r.db.Create(txn).Error
}

func (r *transactionRepo) GetByWallet(wid uint, limit int) ([]models.Transaction, error) {
    var txs []models.Transaction
    err := r.db.Where("wallet_id = ?", wid).Order("created_at desc").Limit(limit).Find(&txs).Error
    return txs, err
}