package handler

import (
	"app/model"
	"app/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Update_One(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var updatedStudent model.Student
		if err := c.ShouldBindJSON(&updatedStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		uc := usecases.NewStudentUseCase()

		err = uc.UpdateOne(c.Request.Context(), id, &updatedStudent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		after_update, err := uc.GetOneByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"student info after update info": after_update,
		})
	}
}
