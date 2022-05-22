package Application

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

type SharedResources interface {
	Share()
}

// Handle request closure data
func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		connectToDatabase(&request)
		//request.DB, request.Connection = connectToDatabase()

		return request
	}
}

func (req *Request) Share() {

}

func (req Request) Response(code int, body interface{}) {
	closeConnection(&req)
	req.Context.JSON(code, body)
}

// Initialize new request closure
func NewRequest(c *gin.Context) Request {
	request := req()
	req := request(c)

	return req
}
