package main

import "github.com/gin-gonic/gin"

func (req Request) OK(body interface{}) {
	req.Response(200, body)
}

func (req Request) NoAuth() {
	req.Response(401, gin.H{
		"message": "You are not authorized ",
	})
}
