package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginControllerServer struct {
	name string
}

func NewGinServer(name string) *ginControllerServer {
	return &ginControllerServer{
		name,
	}
}

func (gs ginControllerServer) getBlogs(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"body": "this is the blog body",
	})
}

func (gs ginControllerServer) editBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"body": "this is the blog body",
	})
}

func (gs ginControllerServer) removeBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"body": "this is the blog body",
	})
}

func (gs ginControllerServer) saveBlog(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"body": "this is the blog body",
	})
}
