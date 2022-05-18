package main

import "github.com/gin-gonic/gin"

func main() {
	// Gin framework for working with routes  (Routing , validate request , response , middleware)
	// gorm (connect database , quiers)
	// auth user
	// language
	// ------------------------------------------------------------------------------------------------- //

	app := app()
	application := app()
	application.Gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong Test Application",
		})
	})
	application.Gin.Run(":9999") //http://127.0.0.1:9999/ping

}
