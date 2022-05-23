package Auth

import (
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"go-starter/Application"
	"go-starter/Models"
)

func CreateUser(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	if !r.IsAdmin {
		r.NoAuth()
		return
	}
	user := Models.User{
		Username: "Youssef Ashraf",
		Email:    "Youssef@youssef.com",
		Password: "123456",
	}

	r.DB.Create(&user)
	r.Created(user)
}

func ViewUser(c *gin.Context) {
	r, auth := Application.AuthRequest(c)
	if !auth {
		return
	}
	//r.OK(r.User)
	r.OK(gin.H{
		"message": gotrans.T("hello_world"),
		"lang":    r.Lang,
	})
}
