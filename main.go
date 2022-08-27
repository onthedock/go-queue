package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Default router
	r := gin.Default()
	// Routes group
	v1 := r.Group("/api/v1")
	{
		// v1.GET("/job/:id", getJob)
		v1.POST("/add/:num1/:num2", submitJob)
	}
	r.Run()
}
