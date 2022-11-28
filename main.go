package main

import (
	"github.com/gin-gonic/gin"
	"github.com/refoscodoc/go-gin-grom-api/controllers"
	"github.com/refoscodoc/go-gin-grom-api/services"
)

func main() {
	r := gin.Default()
	services.ConnectDatabase()

	r.GET("/guitars", controllers.FindGuitars)
	r.POST("/guitars", controllers.CreateGuitar)
	// r.GET("/guitars/:id", controllers.FindGuitar)
	// r.PATCH("/guitars/:id", controllers.UpdateGuitar)
	// r.DELETE("/guitars/:id", controllers.DeleteGuitar)

	err := r.Run()
	if err != nil {
		return
	}
}
