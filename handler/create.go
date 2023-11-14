package handler

import (
	"app/payload"
	"app/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var Data = payload.AddStudentRequest{}
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())

		if err := c.ShouldBind(&Data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := validate.Struct(Data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		student := Data.ToModel()

		uc := usecases.NewStudentUseCase()

		// Pass the context and the address of the student (pointer to model.Student)
		err = uc.InsertOne(c.Request.Context(), student)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": student.ID,
		})
	}
}
