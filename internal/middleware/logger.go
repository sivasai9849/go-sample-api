package middleware

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/sivasai9849/go-advanced-api/pkg/logger"
)

func LoggerMiddleware(logger *logger.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        method := c.Request.Method

        c.Next()

        duration := time.Since(start)
        statusCode := c.Writer.Status()

        logger.Info("| %d | %s | %s | %v |",
            statusCode,
            method,
            path,
            duration,
        )
    }
}