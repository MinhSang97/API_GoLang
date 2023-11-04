package handler

import (
	"app/model"
	"app/payload"
	"app/repo/mysql"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

//func (students payload.AddStudentRequest) TableName() string {
//	return students
//}

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
		}

		student := Data.ToModel()

		repo := mysql.NewStudentRepository(db)

		err_InsertOne := repo.InsertOne(context.Background(), student)
		if err != nil {
			fmt.Println(err_InsertOne)
			os.Exit(1)
		}
		fmt.Println(student)

		c.JSON(http.StatusOK, gin.H{
			"data": student.ID,
		})

	}
}

func GetIdStudent(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		repo := mysql.NewStudentRepository(db)

		id_student, err := repo.GetOneByID(context.Background(), id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		c.JSON(http.StatusOK, gin.H{
			"student": id_student,
		})
	}
}

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

		repo := mysql.NewStudentRepository(db)

		err = repo.UpdateOne(context.Background(), id, &updatedStudent)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		after_update, err := repo.GetOneByID(context.Background(), id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		c.JSON(http.StatusOK, gin.H{
			"student info after update info": after_update,
		})
	}
}
