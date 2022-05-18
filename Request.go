package main

import "github.com/gin-gonic/gin"

type Request struct {
	Context *gin.Context
}

func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		return request

	}
}

func newRequest(c *gin.Context) Request {
	request := req()
	req := request(c)

	return req
}
