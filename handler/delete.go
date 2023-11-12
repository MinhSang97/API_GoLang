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

func Delete_One(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		repo := mysql.NewStudentRepository(db)

		err = repo.DeleteOne(context.Background(), id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		c.JSON(http.StatusOK, gin.H{
			"Delete student by id": id,
		})

	}
}
