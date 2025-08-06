package service

import (
    "errors"

    "github.com/narwar-veer/teen-wallet-api/internal/models"
    "github.com/narwar-veer/teen-wallet-api/internal/repository"
)


type WalletService struct {
    wallets repository.WalletRepository
    txRepo  repository.TransactionRepository
}

func NewWalletService(w repository.WalletRepository, t repository.TransactionRepository) *WalletService {
    return &WalletService{wallets: w, txRepo: t}
}

func (s *WalletService) Deposit(uid uint, amount int64, desc string) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }
    w, err := s.wallets.GetByUserID(uid)
    if err != nil {
        return err
    }
    w.Balance += amount
    if err := s.wallets.Update(w); err != nil {
        return err
    }
    return s.txRepo.Create(&models.Transaction{WalletID: w.ID, Amount: amount, Type: models.Deposit, Description: desc})
}

func (s *WalletService) Withdraw(uid uint, amount int64, desc string) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }
    w, err := s.wallets.GetByUserID(uid)
    if err != nil {
        return err
    }
    if w.Balance < amount {
        return errors.New("insufficient funds")
    }
    w.Balance -= amount
    if err := s.wallets.Update(w); err != nil {
        return err
    }
    return s.txRepo.Create(&models.Transaction{WalletID: w.ID, Amount: -amount, Type: models.Withdraw, Description: desc})
}


func (s *WalletService) Transfer(fromID, toID int, amount int64) error {
    if fromID == toID {
        return errors.New("cannot transfer to self")
    }

    if amount <= 0 {
        return errors.New("amount must be positive")
    }

    return s.wallets.TransferFunds(fromID, toID, amount)
}


func (s *WalletService) Balance(uid uint) (int64, error) {
    w, err := s.wallets.GetByUserID(uid)
    if err != nil {
        return 0, err
    }
    return w.Balance, nil
}