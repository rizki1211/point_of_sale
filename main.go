package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizki1211/point_of_sale/controllers/tokocontroller"
	"github.com/rizki1211/point_of_sale/controllers/userscontroller"
	"github.com/rizki1211/point_of_sale/models"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/toko", tokocontroller.Index)
	r.GET("/toko/:id", tokocontroller.Show)
	r.POST("/toko", tokocontroller.Store)
	r.PUT("/toko/:id", tokocontroller.Update)
	r.DELETE("/toko", tokocontroller.Destroy)

	r.GET("/users", userscontroller.Index)
	r.GET("/users/:id", userscontroller.Show)
	r.POST("/users", userscontroller.Store)
	r.PUT("/users/:id", userscontroller.Update)
	r.DELETE("/users", userscontroller.Destroy)

	r.Run()
}