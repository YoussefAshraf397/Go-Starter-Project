package Visitors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"go-starter/Application"
	"go-starter/Models"
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
	err := validation.Errors{
		"username": validation.Validate(user.Username, validation.Required, validation.Length(2, 20)),
		"email":    validation.Validate(user.Email, validation.Required, validation.Length(5, 500)),
		"password": validation.Validate(user.Password, validation.Required, validation.Length(8, 15)),
	}.Filter()

	if err != nil {
		r.BadRequest(err)
		return
	}
	r.OK(user)

	fmt.Println(err)

}
