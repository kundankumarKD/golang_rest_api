package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"product-api/internal/config"
	httpDelivery "product-api/internal/delivery/http"
	"product-api/internal/domain"
	"product-api/internal/repository"
	"product-api/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Database Connection
	db, err := gorm.Open(sqlite.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-Migrate Schema
	if err := db.AutoMigrate(&domain.User{}, &domain.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	// Initialize Handlers
	authHandler := httpDelivery.NewAuthHandler(userRepo, cfg)
	productHandler := httpDelivery.NewProductHandler(productRepo)

	// Setup Router
	r := gin.Default()
	httpDelivery.SetupRoutes(r, authHandler, productHandler, cfg)

	// Initialize Logger
	logger.InitLogger()

	// Start Server
	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	go func() {
		logger.Log.Info("Starting server", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Error("Failed to start server", "error", err)
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Error("Server forced to shutdown", "error", err)
		log.Fatal("Server forced to shutdown:", err)
	}

	logger.Log.Info("Server exiting")
}
