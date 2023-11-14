package handler

import (
	"app/payload"
	"app/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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

		uc := usecases.NewStudentUseCase()

		studentall, err := uc.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"studentall": studentall, // Change this to "studentall"
		})
	}
}

//func ListStudent(db *gorm.DB) func(*gin.Context) {
//	return func(c *gin.Context) {
//
//		var Data = payload.Paging{}
//
//		if err := c.ShouldBind(&Data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		repo := mysql.NewStudentRepository(db)
//
//		studentsall, err := repo.GetAll(context.Background())
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//		fmt.Println(studentsall)
//
//		c.JSON(http.StatusOK, gin.H{
//			"studentall": studentsall,
//		})
//	}
//}
