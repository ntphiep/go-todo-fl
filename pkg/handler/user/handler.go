package user

import (
	"log"
	"net/http"

	"strconv"

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

func GetUserList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// list users

		var users []data.User

		if err := db.Find(&users).Error; err != nil {
			log.Println("Cannot list users:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot list users"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func GetUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get an user by ID

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var user data.User
		if err := db.First(&user, id).Error; err != nil {
			log.Println("Cannot get user by ID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get user by ID"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func UpdateUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// update an user by ID

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var user data.User
		if err := db.First(&user, id).Error; err != nil {
			log.Println("Cannot get user by ID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get user by ID"})
			return
		}

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Save(&user).Error; err != nil {
			log.Println("Cannot update user by ID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update user by ID"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func DeleteUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// delete an user by ID

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var user data.User
		if err := db.First(&user, id).Error; err != nil {
			log.Println("Cannot get user by ID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get user by ID"})
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			log.Println("Cannot delete user by ID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete user by ID"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
