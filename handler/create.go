package handler

import (
	"app/payload"
	"app/repo/mysql"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"os"
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

		repo := mysql.NewStudentRepository(db)

		err_insertOne := repo.InsertOne(context.Background(), student)
		if err_insertOne != nil {
			fmt.Println(err_insertOne)
			os.Exit(1)
		}
		fmt.Println(student)

		c.JSON(http.StatusOK, gin.H{
			"data": student.ID,
		})
	}
}
