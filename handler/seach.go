package handler

import (
	"app/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SearchStudents(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// Lấy giá trị của các tham số từ query string
		Value := c.Query("Value")
		//FirstName := c.Query("FirstName")

		fmt.Println("Seach with like", Value)

		// Kiểm tra xem có ít nhất một tham số được truyền vào không
		if Value == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "At least one search parameter is required",
			})
			return
		}

		uc := usecases.NewStudentUseCase()

		students, err := uc.Search(c.Request.Context(), Value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"students": students, // Update the key to match the actual data structure
		})
	}
}
