package todo

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ntphiep/go-todo-pg/pkg/data"
	"gorm.io/gorm"
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

		if dataItem.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
			return
		}

		dataItem.Status = "Doing"
		if err := db.Create(&dataItem).Error; err != nil {
			log.Println("Cannot create item:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem})
	}
}

func GetListOfItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// list items

		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}

		var paging DataPaging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []data.ToDoItem

		if err := db.Table(data.ToDoItem{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func ReadItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get an item by ID

		var dataItem data.ToDoItem

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&dataItem).Error; err != nil {
			log.Println("Cannot find item:", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Cannot find item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem})
	}
}

func EditItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// edit an item by ID

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem data.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func DeleteItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// delete an item by ID

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(data.ToDoItem{}.TableName()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
