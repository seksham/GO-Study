package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents the user model
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// users is our mock database
var users = make(map[string]User)

func main() {
	// Create default gin router
	r := gin.Default()

	// Basic route returns hello world
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Group user routes
	userRoutes := r.Group("/users")
	{
		// Get all users
		userRoutes.GET("", func(c *gin.Context) {
			// Convert map to slice for response
			userList := make([]User, 0, len(users))
			for _, user := range users {
				userList = append(userList, user)
			}
			c.JSON(http.StatusOK, userList)
		})

		// Get user by ID
		userRoutes.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if user, exists := users[id]; exists {
				c.JSON(http.StatusOK, user)
				return
			}
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		})

		// Create new user
		userRoutes.POST("", func(c *gin.Context) {
			var newUser User
			if err := c.ShouldBindJSON(&newUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Generate simple ID (in practice, use UUID)
			newUser.ID = string(len(users) + 1)
			users[newUser.ID] = newUser

			c.JSON(http.StatusCreated, newUser)
		})

		// Update user
		userRoutes.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if _, exists := users[id]; !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}

			var updatedUser User
			if err := c.ShouldBindJSON(&updatedUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			updatedUser.ID = id
			users[id] = updatedUser

			c.JSON(http.StatusOK, updatedUser)
		})

		// Delete user
		userRoutes.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if _, exists := users[id]; !exists {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}

			delete(users, id)
			c.Status(http.StatusNoContent)
		})
	}

	// Query parameter example
	r.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "default search")
		c.JSON(http.StatusOK, gin.H{
			"search_query": query,
		})
	})

	// Form parameter example
	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"email": email,
		})
	})

	// Run the server
	r.Run(":8080")
}
