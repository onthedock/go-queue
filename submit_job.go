package main

import (
	"net/http"
	"queue/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
