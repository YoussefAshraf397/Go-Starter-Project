package Visitors

import (
	"github.com/gin-gonic/gin"
	"go-starter/Application"
	"go-starter/Models"
	"go-starter/Validation/Visitors"
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

//func ViewUser(c *gin.Context) {
//	r := Application.NewRequestWithAuth(c)
//	if !r.IsAuth {
//		r.NoAuth()
//		return
//	}
//	r.OK(r.User)
//}

func Register(c *gin.Context) {
	// Binding Request
	r := Application.NewRequest(c)
	var user Models.User
	r.Context.ShouldBind(&user)

	//Validate Request
	r.ValidateRequest(Visitors.RegisterValidation(user))
	if r.ValidationError != nil {
		r.BadRequest(r.ValidationError)
		return
	}
	r.OK(user)
}
