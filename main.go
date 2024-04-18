package main

import (
	"github.com/gin-gonic/gin"
	userscontroller "github.com/rizki1211/point_of_sale/controllers/usersController"
	"github.com/rizki1211/point_of_sale/models"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/users", userscontroller.Index)
	r.GET("/users/:id", userscontroller.Show)
	r.POST("/users", userscontroller.Store)
	r.PUT("/users/:id", userscontroller.Update)
	r.DELETE("/users", userscontroller.Destroy)

	r.Run()
}