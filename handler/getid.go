package handler

import (
	"app/repo/mysql"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
)

func GetId(db *gorm.DB) func(*gin.Context) {
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
