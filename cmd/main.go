package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ntphiep/go-todo-pg/pkg/handler/todo"
)

func main() {
	dsn := "root:my-root-pass@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}


	log.Println("Connected:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/items", todo.CreateItem(db))           // create item
		v1.GET("/items", todo.GetListOfItems(db))        // list items
		v1.GET("/items/:id", todo.ReadItemById(db))      // get an item by ID
		v1.PUT("/items/:id", todo.EditItemById(db))      // edit an item by ID
		v1.DELETE("/items/:id", todo.DeleteItemById(db)) // delete an item by ID
	}

	router.Run(":8080")
}
