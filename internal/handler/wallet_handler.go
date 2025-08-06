package handler

import (
    "net/http"
    "strconv"
    "fmt"
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
type transferRequest struct {
    Amount int64 `json:"amount" binding:"required,gt=0"`
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
     senderID := int(c.GetUint("uid"))
    toUserIDStr := c.Param("to")

    toUserID, err := strconv.Atoi(toUserIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipient ID"})
        return
    }

    var req transferRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        fmt.Println("BindJSON error:", err) // 👈 debug log
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
        return
    }

    err = h.wallet.Transfer(senderID, toUserID, req.Amount)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
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
