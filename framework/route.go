package framework

import (
	"app/dbutil"
	"app/handler"
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
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", handler.CreateItem(db))
			items.GET("")
			items.GET("/:id", handler.GetIdStudent(db))
			items.PATCH("/:id", handler.Update_One(db))
			items.DELETE("/:id")

		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
