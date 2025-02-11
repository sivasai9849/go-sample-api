package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sivasai9849/go-advanced-api/internal/domain"
	"github.com/sivasai9849/go-advanced-api/internal/dto"
	"github.com/sivasai9849/go-advanced-api/internal/service"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

func (h *UserHandler) Create(c *gin.Context) {
    _, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
    defer cancel()

    var input dto.CreateUserRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.Error(domain.ErrInvalidInput)
        return
    }

    user := &domain.User{
        ID:       uuid.New(),
        Email:    input.Email,
        Password: input.Password,
        Name:     input.Name,
    }

    if err := h.userService.Create(user); err != nil {
        switch err {
        case domain.ErrUserExists:
            c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
        default:
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) List(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
        return
    }

    limit, err := strconv.Atoi(limitStr)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
        return
    }
    users, err := h.userService.List(page, limit)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Get(c *gin.Context) {
    id := c.Param("id")
    user, err := h.userService.GetByID(id)
    if err != nil {
        if err == domain.ErrUserNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(c *gin.Context) {
    id := c.Param("id")
    
    // Parse UUID
    userID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
        return
    }

    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
        Name     string `json:"name" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := &domain.User{
        ID:       userID,
        Email:    input.Email,
        Password: input.Password,
        Name:     input.Name,
    }

    if err := h.userService.Update(user); err != nil {
        switch err {
        case domain.ErrUserNotFound:
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        case domain.ErrUserExists:
            c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
        default:
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    if err := h.userService.Delete(id); err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}