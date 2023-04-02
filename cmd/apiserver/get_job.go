package main

import (
	"net/http"
	"queue/jobs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getJob(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "jobId is required",
		})
		return
	}

	juuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "unable to convert jobId: " + c.Param("id"),
		})
		return
	}

	job, err := jobs.Load(juuid.String() + ".json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unexpected error",
		})
		return
	}
	c.JSON(http.StatusOK, job)
}
