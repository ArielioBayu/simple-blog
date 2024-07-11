package main

import (
	"go-simple-blog/config"
	filmcontroller "go-simple-blog/controllers/productcontroller"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	config.ConnectDB()

	r.GET("/films", filmcontroller.Index)
	r.GET("/films/:id", filmcontroller.Show)
	r.POST("/films", filmcontroller.Create)
	r.PUT("/films/:id", filmcontroller.Update)
	r.DELETE("/films", filmcontroller.Delete)

	r.Run()
}
