package repository

import (
    "fmt"

    "github.com/narwar-veer/teen-wallet-api/internal/models"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type WalletRepository interface {
    Create(w *models.Wallet) error
    GetByUserID(uid uint) (*models.Wallet, error)
    Update(w *models.Wallet) error
    TransferFunds(fromID, toID int, amount int64) error
}

type walletRepo struct {
    db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
    return &walletRepo{db: db}
}

func (r *walletRepo) Create(w *models.Wallet) error {
    return r.db.Create(w).Error
}

func (r *walletRepo) GetByUserID(uid uint) (*models.Wallet, error) {
    var w models.Wallet
    if err := r.db.Where("user_id = ?", uid).First(&w).Error; err != nil {
        return nil, err
    }
    return &w, nil
}

func (r *walletRepo) Update(w *models.Wallet) error {
    return r.db.Save(w).Error
}

func (r *walletRepo) TransferFunds(fromID, toID int, amount int64) error {
    tx := r.db.Begin()
    if tx.Error != nil {
        return tx.Error
    }

    var fromWallet, toWallet models.Wallet

    // Lock sender wallet
    if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
        Where("user_id = ?", fromID).
        First(&fromWallet).Error; err != nil {
        tx.Rollback()
        return fmt.Errorf("sender not found")
    }

    // Lock recipient wallet
    if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
        Where("user_id = ?", toID).
        First(&toWallet).Error; err != nil {
        tx.Rollback()
        return fmt.Errorf("recipient not found")
    }

    // Check sufficient funds
    if fromWallet.Balance < amount {
        tx.Rollback()
        return fmt.Errorf("insufficient balance")
    }

    // Perform balance update
    fromWallet.Balance -= amount
    toWallet.Balance += amount

    if err := tx.Save(&fromWallet).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Save(&toWallet).Error; err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit().Error
}
