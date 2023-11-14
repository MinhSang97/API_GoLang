package handler

import (
	"app/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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

		//repo := mysql.NewStudentRepository(db)
		//
		//err = repo.DeleteOne(context.Background(), id)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(1)
		//}

		uc := usecases.NewStudentUseCase()

		err = uc.DeleteOne(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Delete student by id": id,
		})

	}
}
