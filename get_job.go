package main

import (
	"net/http"
	"queue/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
