package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ntphiep/go-todo-pg/pkg/data"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem data.TodoItemCreate

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// preprocess title - trim all spaces
		if err := db.Create(&dataItem).Error; err != nil {
			log.Println("Cannot create item:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Cannot create item hmm",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": dataItem.Id,
		})
	}
}

func GetListOfItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []data.ToDoItem
		var paging data.DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		db = db.Where("status <> ?", "deleted")

		if err := db.Table(data.ToDoItem{}.TableName()).
			Count(&paging.Total).
			Offset((paging.Page - 1) * paging.Limit).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   result,
			"paging": paging,
		})
	}
}

func GetItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem data.ToDoItem

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// find item by ID
		if err := db.
			Where("id = ?", id).
			First(&dataItem).Error; err != nil {

			log.Println("Cannot find item:", err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Cannot find item",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": dataItem,
		})
	}
}

func EditItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem data.TodoItemEdit

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.
			Where("id = ?", id).
			Updates(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func DeleteItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// delete an item by ID

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Table(data.ToDoItem{}.TableName()).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"status": "deleted",
			}).
			Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
