package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User

func Main() {
	router := gin.Default()

	// Endpoint to get all users
	router.GET("/users", findAll)

	// Endpoint to get a specific user by ID
	router.GET("/users/:id", findById)

	// Endpoint to create a new user
	router.POST("/users", save)

	// Endpoint to update a user by ID
	router.PUT("/users/:id", update)

	// Endpoint to delete a user by ID
	router.DELETE("/users/:id", delete)

	// Run the server
	router.Run(":8080")
}

func findAll(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func findById(c *gin.Context) {
	id := c.Param("id")

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func save(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = uuid.New().String() // Generate a new UUID for user ID

	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func update(c *gin.Context) {
	id := c.Param("id")
	var updatedUser User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, user := range users {
		if user.ID == id {
			users[index] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func delete(c *gin.Context) {
	id := c.Param("id")

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
