package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/yourusername/teen-wallet-api/internal/service"
)

type AuthHandler struct {
    auth *service.AuthService
}

func NewAuthHandler(a *service.AuthService) *AuthHandler { return &AuthHandler{auth: a} }

type registerReq struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Phone    string `json:"phone" binding:"required"`
    Password string `json:"password" binding:"required,min=6"`
}

type loginReq struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req registerReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user, err := h.auth.Register(req.Name, req.Email, req.Phone, req.Password)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": user.ID})
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req loginReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    token, err := h.auth.Login(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}
