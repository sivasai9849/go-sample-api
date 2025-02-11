package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sivasai9849/go-advanced-api/internal/domain"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            switch err {
            case domain.ErrUserNotFound:
                c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            case domain.ErrUserExists:
                c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
            case domain.ErrInvalidInput:
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            default:
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
            }
        }
    }
}