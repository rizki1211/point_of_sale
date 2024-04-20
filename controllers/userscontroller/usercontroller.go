package userscontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rizki1211/point_of_sale/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var users []models.Users
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Show(c *gin.Context) {
	var users models.Users
	id := c.Param("id")

	if err := models.DB.First(&users, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Pengguna tidak ditemukan!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": users})
}

func Store(c *gin.Context) {
	var users models.Users

	err := c.ShouldBindJSON(&users)
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error pada field %s dengan kondisi %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

    if !isEmailUnique(users.Email) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah digunakan"})
        return
    }

	if !isUsernameUnique(users.Username) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username sudah digunakan"})
        return
    }

	password := []byte(users.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

    if err != nil {
        panic(err)
    }

	users.Password = string(hashedPassword)

	models.DB.Create(&users)
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil membuat pengguna!"})
}

func Update(c *gin.Context) {
	var users models.Users
	id := c.Param("id")
	
	var userNew struct {
		IdLevel  int64     `json:"id_level"`
		IdToko   int64     `json:"id_toko"`
		Nama     string    `json:"nama"`
		Username string    `json:"username"`
		Password string    `json:"password"`
		Email    string    `json:"email" binding:"email"`
	}

	err := c.ShouldBindJSON(&userNew)
	if err != nil {
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error pada field %s dengan kondisi %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	models.DB.First(&users, id)
	if users.Email != userNew.Email{
		if !isEmailUnique(userNew.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah digunakan"})
			return
		}
	}

	if users.Username != userNew.Username{
		if !isUsernameUnique(userNew.Username) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username sudah digunakan"})
			return
		}
	}

	if userNew.Password != "" && len(userNew.Password) >= 8 {
		bytes, _ := bcrypt.GenerateFromPassword([]byte(userNew.Password), 10)
		userNew.Password = string(bytes)
	} else if userNew.Password != "" && len(userNew.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password harus memiliki setidaknya 8 karakter",
		})
		return
	}

	if models.DB.Model(&users).Where("id = ?", id).Updates(&userNew).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Pengguna tidak berhasil diperbarui"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil diperbarui!"})
}

func Destroy(c *gin.Context) {
	var users models.Users

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&users, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Pengguna tidak berhasil dihapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil dihapus!"})
}


func isEmailUnique(email string) bool {
    var existingUser models.Users
    if err := models.DB.Where("email = ?", email).First(&existingUser).Error; err != nil {
        return true 
    }
    return false 
}

func isUsernameUnique(username string) bool {
    var existingUser models.Users
    if err := models.DB.Where("username = ?", username).First(&existingUser).Error; err != nil {
        return true
    }
    return false
}
