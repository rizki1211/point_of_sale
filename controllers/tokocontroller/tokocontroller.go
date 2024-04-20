package tokocontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rizki1211/point_of_sale/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var toko []models.Toko
	models.DB.Find(&toko)
	c.JSON(http.StatusOK, gin.H{"toko": toko})
}

func Show(c *gin.Context) {
	var toko models.Toko
	id := c.Param("id")

	if err := models.DB.First(&toko, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Toko tidak ditemukan!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"toko": toko})
}

func Store(c *gin.Context) {
	var toko models.Toko

	err := c.ShouldBindJSON(&toko)
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

	 // Check uniqueness of NamaToko
    if !isNamaTokoUnique(toko.NamaToko) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Nama Toko sudah digunakan!"})
        return
    }

	models.DB.Create(&toko)
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil membuat Toko!"})
}

func Update(c *gin.Context) {
	var toko models.Toko

	id := c.Param("id")

	err := c.ShouldBindJSON(&toko)
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

	if models.DB.Model(&toko).Where("id = ?", id).Updates(&toko).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Toko tidak berhasil diperbarui"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Toko berhasil diperbarui!"})
}

func Destroy(c *gin.Context) {
	var toko models.Toko

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&toko, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Toko tidak berhasil dihapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Toko berhasil dihapus!"})
}

func isNamaTokoUnique(namaToko string) bool {
    var existingToko models.Toko
    if err := models.DB.Where("nama_toko = ?", namaToko).First(&existingToko).Error; err != nil {
        return true // NamaToko is unique if not found in database
    }
    return false
}