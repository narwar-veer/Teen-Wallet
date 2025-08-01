package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/narwar-veer/teen-wallet-api/internal/service"
)

type WalletHandler struct {
    wallet *service.WalletService
}

func NewWalletHandler(w *service.WalletService) *WalletHandler { return &WalletHandler{wallet: w} }

type amountReq struct {
    Amount int64  `json:"amount" binding:"required,gt=0"`
    Desc   string `json:"description"`
}

func (h *WalletHandler) Deposit(c *gin.Context) {
    uid := c.GetUint("uid")
    var req amountReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.wallet.Deposit(uid, req.Amount, req.Desc); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *WalletHandler) Withdraw(c *gin.Context) {
    uid := c.GetUint("uid")
    var req amountReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.wallet.Withdraw(uid, req.Amount, req.Desc); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *WalletHandler) Transfer(c *gin.Context) {
    fromUID := c.GetUint("uid")
    toUIDStr := c.Param("to")
    toUID64, _ := strconv.ParseUint(toUIDStr, 10, 32)
    var req amountReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.wallet.Transfer(fromUID, uint(toUID64), req.Amount); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

// Balance returns wallet balance
func (h *WalletHandler) Balance(c *gin.Context) {
    uid := c.GetUint("uid")
    bal, err := h.wallet.Balance(uid)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"balance": bal})
}
