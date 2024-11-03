package todo

import (
	"log"
	"net/http"
	"strconv"

	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/ntphiep/go-todo-pg/pkg/data"
)


func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// create item

		var dataItem data.ToDoItem
		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// preprocess title - trim all spaces
		dataItem.Title = strings.TrimSpace(dataItem.Title)
	}
}

func GetListOfItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// list items
	}
}

func ReadItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get an item by ID
	}
}

func EditItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// edit an item by ID
	}
}

func DeleteItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// delete an item by ID
	}
}