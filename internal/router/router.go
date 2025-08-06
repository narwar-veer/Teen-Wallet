package router

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/narwar-veer/teen-wallet-api/internal/config"
    "github.com/narwar-veer/teen-wallet-api/internal/handler"
    "github.com/narwar-veer/teen-wallet-api/internal/middleware"
    "github.com/narwar-veer/teen-wallet-api/internal/repository"
    "github.com/narwar-veer/teen-wallet-api/internal/service"
)

func New(cfg *config.Config, db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // repositories
    userRepo := repository.NewUserRepository(db)
    walletRepo := repository.NewWalletRepository(db)
     
    txRepo := repository.NewTransactionRepository(db)

    // services
    authSrv := service.NewAuthService(userRepo, walletRepo, cfg.JWT)
    walletSrv := service.NewWalletService(walletRepo, txRepo)

    // handlers
    authH := handler.NewAuthHandler(authSrv)
    walletH := handler.NewWalletHandler(walletSrv)

    // public routes
    r.POST("/v1/auth/register", authH.Register)
    r.POST("/v1/auth/login", authH.Login)

    // protected routes (JWT)
    authGroup := r.Group("/v1").Use(middleware.AuthMiddleware(authSrv))
    authGroup.POST("/wallet/deposit", walletH.Deposit)
    authGroup.POST("/wallet/withdraw", walletH.Withdraw)
    authGroup.POST("/wallet/transfer/:to", walletH.Transfer)
    authGroup.GET("/wallet/balance", walletH.Balance)

    return r
}