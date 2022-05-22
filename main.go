package main

import (
	"github.com/gin-gonic/gin"
	"go-starter/Application"
	"go-starter/Models"
)

func main() {
	// Gin framework for working with routes  (Routing , validate request , response , middleware)
	// gorm (connect database , quiers)
	// auth user
	// language
	// ------------------------------------------------------------------------------------------------- //

	app := Application.NewApp()
	// migrate project
	app.DB.AutoMigrate(&Models.User{})
	// close connection
	Application.CloseConnection(&app)

	app.Gin.GET("/create-user", func(c *gin.Context) {
		r := Application.NewRequest(c)
		user := Models.User{
			Username: "Youssef Ashraf",
			Email:    "Youssef@youssef.com",
			Password: "123456",
		}

		r.DB.Create(&user)
		r.Created(user)
	})
	app.Gin.Run(":9999") //http://127.0.0.1:9999/ping

}
