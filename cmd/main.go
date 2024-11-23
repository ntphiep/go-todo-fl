package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ntphiep/go-todo-pg/internal/utils"
	"github.com/ntphiep/go-todo-pg/pkg/handler/todo"
	"github.com/ntphiep/go-todo-pg/pkg/handler/user"
)

func main() {
	dsn := "root:my-root-pass@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
	router := gin.Default()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected:", db)

	v1 := router.Group("/v1")

	// todo
	{
		items := v1.Group("/items")
		items.POST("", todo.CreateItem(db))           // create item
		items.GET("", todo.GetListOfItems(db))        // list items
		items.GET("/:id", todo.ReadItemById(db))      // get an item by ID
		items.PUT("/:id", todo.EditItemById(db))      // edit an item by ID
		items.DELETE("/:id", todo.DeleteItemById(db)) // delete an item by ID
	}

	// user
	{
		v1.POST("/users", user.CreateUser(db))           // create user
		v1.GET("/users", user.GetUserList(db))           // list users
		v1.GET("/users/:id", user.GetUserById(db))       // get an user by ID
		v1.PUT("/users/:id", user.UpdateUserById(db))    // edit an user by ID
		v1.DELETE("/users/:id", user.DeleteUserById(db)) // delete an user by ID
	}

	// Enable CORS
	enhancedRouter := utils.EnableCORS(utils.SetJSONContentType(router))

	// router.Run(":8080")\
	log.Fatal(http.ListenAndServe(":8080", enhancedRouter))
}
