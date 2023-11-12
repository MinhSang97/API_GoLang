package framework

import (
	"app/dbutil"
	"app/handler"
	"app/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Route() {
	db := dbutil.ConnectDB()
	fmt.Println("Connected: ", db)

	// CRUD: Create, Read, Update, Delete
	// POST /v1/items (create a new item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// (PUT | PATCH) /v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (delete item by id)
	r := gin.Default()
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.BasicAuthMiddleware())

	r.GET("/secure", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a secure route"})
	})

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", handler.CreateItem(db))
			items.GET("", handler.GetAllStudent(db))
			items.GET("/:id", handler.GetId(db))
			items.PATCH("/:id", handler.Update_One(db))
			items.DELETE("/:id", handler.Delete_One(db))

		}
	}

	r.GET("/ping", func(c *gin.Context) {
		// Gây ra một lỗi để kiểm tra middleware
		panic("Some error occurred")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
