package services

import (
	"net/http"
	"os"

	"user-service/src/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

var jwtSecret = os.Getenv("JWT_SECRET_KEY")

func InitDatabase(database *gorm.DB) {
	db = database
}

// Register a new user
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Authenticate a user
func AuthenticateUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := generateJWT(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Generate JWT
func generateJWT(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
	})
	tokenString, _ := token.SignedString([]byte(jwtSecret))
	return tokenString
}

// List all users (admin only)
func ListUsers(c *gin.Context) {
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Get user details
func GetUserDetails(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update user information
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// Deactivate user account
func DeactivateUser(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Deleted successfully"})
}

// Get user preferences
func GetUserPreferences(c *gin.Context) {
	id := c.Param("id")
	var preferences []models.UserPreference
	db.Where("user_id = ?", id).Find(&preferences)
	c.JSON(http.StatusOK, preferences)
}

// Update user preferences
func UpdateUserPreferences(c *gin.Context) {
	id := c.Param("id")
	var preferences []models.UserPreference
	if err := c.ShouldBindJSON(&preferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, pref := range preferences {
		db.Where("user_id = ? AND key = ?", id, pref.Key).Assign(pref).FirstOrCreate(&pref)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Preferences updated successfully"})
}
