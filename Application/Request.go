package Application

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-starter/Models"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
	IsAuth     bool
	User       Models.User
	IsAdmin    bool
}

type SharedResources interface {
	Share()
}

// Handle request closure data
func req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.Context = c
		connectToDatabase(&request)
		//request.DB, request.Connection = connectToDatabase()

		return &request
	}
}

func (req *Request) Share() {

}

func (req Request) Response(code int, body interface{}) {
	CloseConnection(&req)
	req.Context.JSON(code, body)
}

// Initialize new request closure
func NewRequest(c *gin.Context) *Request {
	request := req()
	req := request(c)

	return req
}

func NewRequestWithAuth(c *gin.Context) *Request {
	return NewRequest(c).Auth()

}

func (req Request) Auth() *Request {

	req.IsAuth = false
	req.IsAdmin = false
	authHeader := req.Context.GetHeader("Authorized")
	if authHeader != "" {
		fmt.Println("Header: ", authHeader)
		req.DB.Where("token = ?", authHeader).First(&req.User)
		if req.User.ID != 0 {
			req.IsAuth = true
		}
		if req.User.Group == "admin" {
			req.IsAdmin = true
		}
	}
	return &req
}
