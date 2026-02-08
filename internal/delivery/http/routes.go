package http

import (
	"product-api/internal/config"
	"product-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler *AuthHandler, productHandler *ProductHandler, cfg config.Config) {
	// Global Middleware
	r.Use(middleware.RequestLogger())
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimitMiddleware())

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		products := api.Group("/products")
		{
			products.GET("/", productHandler.GetAllProducts)
			products.GET("/:id", productHandler.GetProductByID)

			// Protected routes
			protected := products.Group("/")
			protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
			{
				protected.POST("/", productHandler.CreateProduct)
				protected.PUT("/:id", productHandler.UpdateProduct)
				protected.DELETE("/:id", productHandler.DeleteProduct)
			}
		}
	}
}
