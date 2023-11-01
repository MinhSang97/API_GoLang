package handler

import (
	"app/payload"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

// func (students model.Student) TableName() string { return "students" }

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data payload.AddStudentRequest
		
		var validate *validator.Validate

		validate = validator.New(validator.WithRequiredStructEnabled())

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		err := validate.Struct(data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		student := data.ToModel()

		if err := db.Create(&student).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": student.ID,
		})
	}
}
