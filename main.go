package main

import (
	"crud/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {

	// init db with gorm and postgres
	dsn := "host=localhost user=postgres password=postgres dbname=golang_crud port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// debugging
	db = db.Debug()

	// init gin
	r := gin.Default()

	// define a simple route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong ok",
		})
	})

	v1 := r.Group("/v1")
	{
		customers := v1.Group("/customers")
		customers.POST("", internal.CreateAPI(db))
		customers.GET("/:id", internal.ViewAPI(db))
		customers.GET("", internal.ListingAPI(db))
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default to 8080 instead of 3000
	}
	fmt.Println("Listening on port " + port)

	_ = r.Run(":" + port)

}
