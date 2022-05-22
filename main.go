package main

import (
	"github.com/gin-gonic/gin"
	"go-starter/Application"
)

func main() {
	// Gin framework for working with routes  (Routing , validate request , response , middleware)
	// gorm (connect database , quiers)
	// auth user
	// language
	// ------------------------------------------------------------------------------------------------- //

	app := Application.NewApp()

	app.Gin.GET("/ping", func(c *gin.Context) {
		r := Application.NewRequest(c)
		r.OK(gin.H{
			"message": "YOUSSEF",
		})
	})
	app.Gin.Run(":9999") //http://127.0.0.1:9999/ping

}
