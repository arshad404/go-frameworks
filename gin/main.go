package main

import (
	"github.com/gin-gonic/gin"
)

// medium app
func main() {
	router := gin.Default()

	controllerServer := NewGinServer("go-gin-controller-server")

	router.POST("/blog", controllerServer.saveBlog)
	router.GET("/blog/:id", controllerServer.getBlogs)
	router.DELETE("/blog/:id", controllerServer.removeBlog)
	router.PUT("/blog", controllerServer.editBlog)

	router.Run() // localhost:8080
}
