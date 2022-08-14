package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default router
	r := gin.Default()
	// Routes group
	v1 := r.Group("/api/v1")
	{
		v1.GET("/job/:id", getJob)
		v1.POST("/add/:num1/:num2", submitJob)
		v1.GET("/result/:result", getResult)
	}
	r.Run()
}

func getJob(c *gin.Context) {
	if c.Param("id") != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "PROCESSING " + c.Param("id"),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "jobId is required",
		})
	}
}

func submitJob(c *gin.Context) {
	if c.Param("num1") != "" && c.Param("num2") != "" {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "processing " + c.Param("num1") + " + " + c.Param("num2"),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error accepting job",
		})
	}
}

func getResult(c *gin.Context) {
	if c.Param("result") != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "result for " + c.Param("result") + " pending",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "result is required",
		})
	}
}
