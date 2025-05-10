// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"user-service/src/cmd/services"
	"user-service/src/middlewares"
	"user-service/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var LAUNCHPORT = ":8084"
var DSN = os.Getenv("POSTGRES_DATABASE_DSN")

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	models.Migrate(db)
	services.InitDatabase(db)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Routes
	r.POST("/api/users", services.RegisterUser)
	r.POST("/api/users/authenticate", services.AuthenticateUser)

	r.GET("/api/users", middlewares.AuthMiddleware, services.ListUsers)
	r.GET("/api/users/:id", middlewares.AuthMiddleware, services.GetUserDetails)
	r.PUT("/api/users/:id", middlewares.AuthMiddleware, services.UpdateUser)
	r.DELETE("/api/users/:id", middlewares.AuthMiddleware, services.DeactivateUser)

	r.GET("/api/users/:id/preferences", middlewares.AuthMiddleware, services.GetUserPreferences)
	r.PUT("/api/users/:id/preferences", middlewares.AuthMiddleware, services.UpdateUserPreferences)

	fmt.Printf("User service 📨 started at http://localhost%s\n",
		LAUNCHPORT)

	if err := r.Run(LAUNCHPORT); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
