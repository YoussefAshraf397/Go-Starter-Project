package Visitors

import (
	"github.com/gin-gonic/gin"
	"go-starter/Application"
)

//func CreateUser(c *gin.Context) {
//	r := Application.NewRequestWithAuth(c)
//	if !r.IsAdmin {
//		r.NoAuth()
//		return
//	}
//	user := Models.User{
//		Username: "Youssef Ashraf",
//		Email:    "Youssef@youssef.com",
//		Password: "123456",
//	}
//
//	r.DB.Create(&user)
//	r.Created(user)
//}

func ViewUser(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	if !r.IsAuth {
		r.NoAuth()
		return
	}
	r.OK(r.User)
}
