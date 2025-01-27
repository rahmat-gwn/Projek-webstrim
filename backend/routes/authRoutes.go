package routes

import (
	"backend/models"
	"github.com/gin-gonic/gin"
)

// RegisterUserHandler handles the registration route
func RegisterUserHandler(c *gin.Context) {
	var user models.User
	// Bind JSON input to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Validate if the fields are provided
	if user.Username == "" || user.Email == "" || user.PhoneNumber == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "All fields (username, email, phone number, password) are required"})
		return
	}

	// Save the new user in the database
	if result := models.DB.Create(&user); result.Error != nil {
		c.JSON(500, gin.H{"error": "Error registering user"})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully", "user": user})
}

// LoginUserHandler handles the login route
func LoginUserHandler(c *gin.Context) {
	var user models.User
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Search for the user by email and password
	if result := models.DB.Where("email = ? AND password = ?", loginData.Email, loginData.Password).First(&user); result.Error != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful", "user": user})
}

// SetupRoutes sets up the routes for user registration and login
func SetupRoutes() {
	r := gin.Default()

	r.POST("/register", RegisterUserHandler)
	r.POST("/login", LoginUserHandler)

	// Start the server on port 8080
	r.Run(":8080")
}
