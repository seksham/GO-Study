package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// Product represents a product model with validation
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required,min=3"`
	Price       float64   `json:"price" binding:"required,gt=0"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

// CustomError represents a custom error response
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// AuthMiddleware checks for a simple API key
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != "your-secret-key" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, CustomError{
				Code:    http.StatusUnauthorized,
				Message: "Invalid API key",
			})
			return
		}
		c.Next()
	}
}

// LoggerMiddleware logs request details
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		duration := time.Since(start)
		log.Printf(
			"[%s] %s %s %d %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Writer.Status(),
			duration,
		)
	}
}

// ErrorHandler middleware handles panics
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				c.JSON(http.StatusInternalServerError, CustomError{
					Code:    http.StatusInternalServerError,
					Message: "Internal server error",
				})
			}
		}()
		c.Next()
	}
}

// RateLimiter implements a simple rate limiting middleware
func RateLimiter() gin.HandlerFunc {
	// Simple map to store client request counts
	// In production, use Redis or similar
	clients := make(map[string]int64)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now().Unix()

		// Clean old entries (simple cleanup)
		if now%10 == 0 { // Cleanup every 10 seconds
			for ip, timestamp := range clients {
				if now-timestamp > 60 { // Remove after 60 seconds
					delete(clients, ip)
				}
			}
		}

		// Check if client has made a request in the last second
		if lastRequest, exists := clients[ip]; exists && now-lastRequest < 1 {
			c.JSON(http.StatusTooManyRequests, CustomError{
				Code:    http.StatusTooManyRequests,
				Message: "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		clients[ip] = now
		c.Next()
	}
}

func main() {
	// Create gin router with default middleware
	r := gin.Default()

	// Add custom middleware
	r.Use(ErrorHandler())
	r.Use(LoggerMiddleware())
	r.Use(RateLimiter())

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-API-Key"}
	r.Use(cors.New(config))

	// API versioning
	v1 := r.Group("/api/v1")
	{
		// Public routes
		public := v1.Group("")
		{
			public.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"status": "healthy",
					"time":   time.Now(),
				})
			})
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(AuthMiddleware())
		{
			// File upload
			protected.POST("/upload", func(c *gin.Context) {
				file, err := c.FormFile("file")
				if err != nil {
					c.JSON(http.StatusBadRequest, CustomError{
						Code:    http.StatusBadRequest,
						Message: "No file uploaded",
					})
					return
				}

				// Save file with timestamp
				filename := time.Now().Format("20060102150405") + "-" + file.Filename
				if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
					c.JSON(http.StatusInternalServerError, CustomError{
						Code:    http.StatusInternalServerError,
						Message: "Failed to save file",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"filename": filename,
				})
			})

			// Async handler example
			protected.GET("/async", func(c *gin.Context) {
				// Copy context to use in goroutine
				cCp := c.Copy()

				// Process async
				go func() {
					time.Sleep(5 * time.Second)
					log.Printf("Done processing for %s", cCp.ClientIP())
				}()

				c.JSON(http.StatusOK, gin.H{
					"message": "Processing started",
				})
			})

			// Custom validation example
			protected.POST("/products", func(c *gin.Context) {
				var product Product
				if err := c.ShouldBindJSON(&product); err != nil {
					c.JSON(http.StatusBadRequest, CustomError{
						Code:    http.StatusBadRequest,
						Message: err.Error(),
					})
					return
				}

				// Set creation time
				product.CreatedAt = time.Now()

				// In real app, save to database
				c.JSON(http.StatusCreated, product)
			})
		}
	}

	// Custom 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, CustomError{
			Code:    http.StatusNotFound,
			Message: "Route not found",
		})
	})

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
