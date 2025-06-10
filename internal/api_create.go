package internal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Here you would typically handle the creation of a resource, e.g., a customer.
		var requestBody Customer
		if err := c.ShouldBind(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := requestBody.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&requestBody)

		// Example API handler
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Customer created successfully",
			"data":    requestBody,
		})
	}
}
