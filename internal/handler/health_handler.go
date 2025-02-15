package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
    return &HealthHandler{}
}

func (h *HealthHandler) Check(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}