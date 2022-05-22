package Visitors

import (
	"github.com/gin-gonic/gin"
	"go-starter/Application"
	"go-starter/Models"
)

func CreateUser(c *gin.Context) {
	r := Application.NewRequest(c)
	user := Models.User{
		Username: "Youssef Ashraf",
		Email:    "Youssef@youssef.com",
		Password: "123456",
	}

	r.DB.Create(&user)
	r.Created(user)
}
