package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Node struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var nodes = []Node{} // In-memory storage

func CreateNode(c *gin.Context) {
	var newNode Node
	if err := c.ShouldBindJSON(&newNode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nodes = append(nodes, newNode)
	c.JSON(http.StatusCreated, newNode)
}

func GetNode(c *gin.Context) {
	id := c.Param("id")
	for _, node := range nodes {
		if node.ID == id {
			c.JSON(http.StatusOK, node)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
}

func ListNodes(c *gin.Context) {
	c.JSON(http.StatusOK, nodes)
}
