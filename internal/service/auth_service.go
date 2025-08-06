package service

import (
    "errors"
    "time"

    "github.com/go-playground/validator/v10"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"

    "github.com/narwar-veer/teen-wallet-api/internal/config"
    "github.com/narwar-veer/teen-wallet-api/internal/models"
    "github.com/narwar-veer/teen-wallet-api/internal/repository"
)


type AuthService struct {
    users  repository.UserRepository
    wallets repository.WalletRepository
    jwtCfg config.JWT
    validate *validator.Validate
}

type AuthClaims struct {
    UID uint `json:"uid"`
    jwt.RegisteredClaims
}

func NewAuthService(u repository.UserRepository, w repository.WalletRepository, jwtCfg config.JWT) *AuthService {
    return &AuthService{users: u, wallets: w, jwtCfg: jwtCfg, validate: validator.New()}
}

func hashPassword(pw string) (string, error) {
    b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
    return string(b), err
}

func verifyPassword(hash, pw string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}

func (s *AuthService) Register(name, email, phone, password string) (*models.User, error) {
    if err := s.validate.Var(email, "required,email"); err != nil {
        return nil, err
    }
    if err := s.validate.Var(password, "min=6"); err != nil {
        return nil, err
    }
    if _, err := s.users.GetByEmail(email); err == nil {
        return nil, errors.New("email already exists")
    }

    hash, err := hashPassword(password)
    if err != nil {
        return nil, err
    }

    u := &models.User{Name: name, Email: email, Phone: phone, PasswordHash: hash}
    if err := s.users.Create(u); err != nil {
        return nil, err
    }
    if err := s.wallets.Create(&models.Wallet{UserID: u.ID, Balance: 0}); err != nil {
    return nil, err
}

return u, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
    u, err := s.users.GetByEmail(email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }
    if !verifyPassword(u.PasswordHash, password) {
        return "", errors.New("invalid credentials")
    }
    claims := AuthClaims{
        UID: u.ID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtCfg.Secret))
}

func (s *AuthService) ParseToken(t string) (*AuthClaims, error) {
    token, err := jwt.ParseWithClaims(t, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(s.jwtCfg.Secret), nil
    })
    if err != nil || !token.Valid {
        return nil, errors.New("invalid token")
    }
    claims, ok := token.Claims.(*AuthClaims)
    if !ok {
        return nil, errors.New("invalid token claims")
    }
    return claims, nil
}