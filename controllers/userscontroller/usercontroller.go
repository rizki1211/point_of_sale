package userscontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizki1211/point_of_sale/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var users []models.Users
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func Show(c *gin.Context){
	var users models.Users
	id := c.Param("id")


	if err := models.DB.First(&users, id).Error; err != nil {
		switch  err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"User not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"users":users})
}

func Store(c *gin.Context){
	var users models.Users

	if err := c.ShouldBindJSON(&users); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
			return
	}

	models.DB.Create(&users)
	c.JSON(http.StatusOK, gin.H{"message":"users success created!"})
}

func Update(c *gin.Context){
	var users models.Users

	id := c.Param("id")

	if err := c.ShouldBindJSON(&users); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
			return
	}

	if models.DB.Model(&users).Where("id = ?", id).Updates(&users).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"user not updated"})
			return	
	}

	c.JSON(http.StatusOK, gin.H{"message":"users success updated!"})
}

func Destroy(c *gin.Context){
	var users models.Users

	var input struct {
		Id json.Number
	}
	 
	if err := c.ShouldBindJSON(&input); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
			return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&users, id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"user not deleted"})
			return
	}

	c.JSON(http.StatusOK, gin.H{"message":"users success deleted!"})
}