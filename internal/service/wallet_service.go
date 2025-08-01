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

// Transfer between teens
func (s *WalletService) Transfer(fromUID, toUID uint, amount int64) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }
    fromW, err := s.wallets.GetByUserID(fromUID)
    if err != nil { return err }
    toW, err := s.wallets.GetByUserID(toUID)
    if err != nil { return err }
    if fromW.Balance < amount { return errors.New("insufficient funds") }

    fromW.Balance -= amount
    toW.Balance += amount
    if err := s.wallets.Update(fromW); err != nil { return err }
    if err := s.wallets.Update(toW); err != nil { return err }

    if err := s.txRepo.Create(&models.Transaction{WalletID: fromW.ID, Amount: -amount, Type: models.TransferOut, Description: "transfer to user"}); err != nil {
        return err
    }
    return s.txRepo.Create(&models.Transaction{WalletID: toW.ID, Amount: amount, Type: models.TransferIn, Description: "transfer from user"})
}

// Balance returns current wallet balance
func (s *WalletService) Balance(uid uint) (int64, error) {
    w, err := s.wallets.GetByUserID(uid)
    if err != nil {
        return 0, err
    }
    return w.Balance, nil
}