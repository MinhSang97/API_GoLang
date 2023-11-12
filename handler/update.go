package handler

import (
	"app/model"
	"app/repo/mysql"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
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
