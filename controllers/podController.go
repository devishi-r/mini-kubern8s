package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pod struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Node string `json:"node"`
}

var pods = []Pod{} // In-memory storage

func CreatePod(c *gin.Context) {
	var newPod Pod
	if err := c.ShouldBindJSON(&newPod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pods = append(pods, newPod)
	c.JSON(http.StatusCreated, newPod)
}

func GetPod(c *gin.Context) {
	id := c.Param("id")
	for _, pod := range pods {
		if pod.ID == id {
			c.JSON(http.StatusOK, pod)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Pod not found"})
}

func DeletePod(c *gin.Context) {
	id := c.Param("id")
	for i, pod := range pods {
		if pod.ID == id {
			pods = append(pods[:i], pods[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Pod deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Pod not found"})
}
