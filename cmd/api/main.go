package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sivasai9849/go-advanced-api/internal/config"
	"github.com/sivasai9849/go-advanced-api/internal/domain"
	"github.com/sivasai9849/go-advanced-api/internal/handler"
	"github.com/sivasai9849/go-advanced-api/internal/middleware"
	postgresRepo "github.com/sivasai9849/go-advanced-api/internal/repository/postgres"
	"github.com/sivasai9849/go-advanced-api/internal/service"
	"github.com/sivasai9849/go-advanced-api/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
    config config.Config
    logger *logger.Logger
    db     *gorm.DB
    router *gin.Engine
}

func NewApp(cfg config.Config, logger *logger.Logger) (*App, error) {
    db, err := initDB(cfg)
    if err != nil {
        return nil, fmt.Errorf("failed to init db: %w", err)
    }

    router := gin.New()
    router.Use(
        gin.Recovery(),
        middleware.LoggerMiddleware(logger),
        middleware.ErrorHandler(),
    )

    return &App{
        config: cfg,
        logger: logger,
        db:     db,
        router: router,
    }, nil
}

func initDB(cfg config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
        cfg.DBHost,
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBName,
        cfg.DBPort,
    )
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    if err := db.AutoMigrate(&domain.User{}); err != nil {
        return nil, fmt.Errorf("failed to migrate database: %w", err)
    }

    return db, nil
}

func (a *App) setupRoutes() {
    userRepo := postgresRepo.NewUserRepository(a.db)

    userService := service.NewUserService(userRepo)

    userHandler := handler.NewUserHandler(userService)

	healthHandler := handler.NewHealthHandler()

    v1 := a.router.Group("/api/v1")
    {
        users := v1.Group("/users")
        {
            users.POST("/", userHandler.Create)
            users.GET("/:id", userHandler.Get)
            users.PUT("/:id", userHandler.Update)
            users.DELETE("/:id", userHandler.Delete)
            users.GET("/", userHandler.List)
        }
    }

    a.router.GET("/health", healthHandler.Check)
}

func (a *App) Run() error {
    a.logger.Info("Setting up routes...")
    a.setupRoutes()
    serverAddr := fmt.Sprintf(":%d", a.config.ServerPort)
    a.logger.Info("%s", "Starting server on " + serverAddr)
    
    return a.router.Run(serverAddr)
}

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    logger := logger.NewLogger()

    app, err := NewApp(cfg, logger)
    if err != nil {
        logger.Fatal("Failed to initialize app: %v", err)
    }

    if err := app.Run(); err != nil {
        logger.Fatal("Error running app: %v", err)
    }
}