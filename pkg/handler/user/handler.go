package user

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ntphiep/go-todo-pg/pkg/data"
	"gorm.io/gorm"
)


func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// creaate user

		var dataUser data.User
		if err := c.ShouldBind(&dataUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if dataUser.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
			return
		}

		if err := db.Create(&dataUser).Error; err != nil {
			log.Println("Cannot create user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataUser})
	}
}