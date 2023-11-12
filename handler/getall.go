package handler

import (
	"app/payload"
	"app/repo/mysql"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func GetAllStudent(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var Data = payload.AddStudentRequest{}

		if err := c.ShouldBind(&Data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		repo := mysql.NewStudentRepository(db)

		studentsall, err := repo.GetAll(context.Background())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(studentsall)

		c.JSON(http.StatusOK, gin.H{
			"studentall": studentsall,
		})
	}
}
