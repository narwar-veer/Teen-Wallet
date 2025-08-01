package repository

import (
    "github.com/narwar-veer/teen-wallet-api/internal/models"
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(user *models.User) error
    GetByEmail(email string) (*models.User, error)
    GetByID(id uint) (*models.User, error)
}

type userRepo struct { db *gorm.DB }

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepo) GetByEmail(email string) (*models.User, error) {
    var u models.User
    if err := r.db.Preload("Wallet").Where("email = ?", email).First(&u).Error; err != nil {
        return nil, err
    }
    return &u, nil
}

func (r *userRepo) GetByID(id uint) (*models.User, error) {
    var u models.User
    if err := r.db.Preload("Wallet").First(&u, id).Error; err != nil {
        return nil, err
    }
    return &u, nil
}