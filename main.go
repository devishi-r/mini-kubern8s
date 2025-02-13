package main

import (
	"minikubernetes/controllers" // Import the controllers package

	"github.com/gin-gonic/gin" // Though not used in this example, it's good practice to keep it if you intend to use it later
)

func main() {
	r := gin.Default()

	// Pod Routes
	r.POST("/pods", controllers.CreatePod)
	r.GET("/pods/:id", controllers.GetPod)
	r.DELETE("/pods/:id", controllers.DeletePod)

	// Node Routes
	r.POST("/nodes", controllers.CreateNode)
	r.GET("/nodes/:id", controllers.GetNode)
	r.GET("/nodes", controllers.ListNodes) //Corrected this line also
	r.Run(":8080")                         // Start the API server
}
