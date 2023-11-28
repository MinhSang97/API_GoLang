package handler

import (
	"app/payload"
	"app/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

//type Paging struct {
//	Page  int   `json:"page" form:"page"`
//	Limit int   `json:"limit" form:"limit"`
//	Total int64 `json:"total" form:"-"`
//}
//
//func (p *Paging) Process() {
//	if p.Page <= 0 {
//		p.Page = 1
//	}
//	if p.Limit <= 0 || p.Limit > 100 {
//		p.Limit = 10
//	}
//}

func GetAllStudent(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var Data = payload.AddStudentRequest{}

		if err := c.ShouldBind(&Data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//var paging = Paging{}
		//paging.Process()
		//
		//if err := db.Count(&paging.Total).Error; err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}

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
