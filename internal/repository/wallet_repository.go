package repository

import (
    "github.com/narwar-veer/teen-wallet-api/internal/models"
    "gorm.io/gorm"
)

type WalletRepository interface {
    Update(wallet *models.Wallet) error
    GetByUserID(uid uint) (*models.Wallet, error)
}

type walletRepo struct { db *gorm.DB }

func NewWalletRepository(db *gorm.DB) WalletRepository {
    return &walletRepo{db: db}
}

func (r *walletRepo) Update(w *models.Wallet) error {
    return r.db.Save(w).Error
}

func (r *walletRepo) GetByUserID(uid uint) (*models.Wallet, error) {
    var w models.Wallet
    if err := r.db.Where("user_id = ?", uid).First(&w).Error; err != nil {
        return nil, err
    }
    return &w, nil
}