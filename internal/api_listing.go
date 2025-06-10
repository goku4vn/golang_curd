package internal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListingAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paginationParams PaginationParams
		if err := c.ShouldBindQuery(&paginationParams); err != nil {
			c.JSON(400, gin.H{"error": "Invalid pagination parameters"})
			return
		}
		// Apply pagination parameters to the query
		offset, limit := paginationParams.GetLimitOffset()
		db = db.Offset(offset).Limit(limit)

		var customers []Customer

		if err := db.Find(&customers).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve customers"})
			return
		}

		dataProvider := NewDataProvider(customers, paginationParams)

		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Customers retrieved successfully",
			"data":    dataProvider,
		})
	}
}
