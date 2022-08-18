package main

import (
	"net/http"
	"queue/models"
	"strconv"

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
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "jobId is required",
		})
		return
	}

	jid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to convert jobId: " + c.Param("id"),
		})
		return
	}

	job, err := models.GetJob(jid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unexpected error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": job,
	})
}

func submitJob(c *gin.Context) {
	if c.Param("num1") == "" || c.Param("num2") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "requires two parameters",
		})
		return
	}

	n1, err1 := strconv.Atoi(c.Param("num1"))
	n2, err2 := strconv.Atoi(c.Param("num2"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to convert paramters",
		})
		return
	}

	result, err := models.AddJob(n1, n2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to process the job",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"operation": "add",
		"num1":      n1,
		"num2":      n2,
		"result":    strconv.Itoa(result),
	})
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
