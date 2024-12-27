package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ginControllerServer struct {
	name        string
	dataStorage dataStorage
}

func NewGinServer(name string) *ginControllerServer {
	return &ginControllerServer{
		name,
		NewDataStorage(),
	}
}

func (gs ginControllerServer) getBlogs(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam) // Convert `id` to integer
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Fetch the blog using the id
	blog := gs.dataStorage.get(id)
	if blog == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	// Return the blog as a JSON response
	ctx.JSON(http.StatusOK, gin.H{
		"body": blog,
	})
}

func (gs ginControllerServer) editBlog(ctx *gin.Context) {
	var requestBody blog

	// Bind the request body to the getRequestBody struct
	if err := ctx.Bind(&requestBody); err != nil {
		// Handle error, e.g., return an error response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := gs.dataStorage.edit(requestBody)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully edited the blog",
	})
}

func (gs ginControllerServer) removeBlog(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam) // Convert `id` to integer
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	fmt.Println(id)
	gs.dataStorage.delete(id)
	ctx.JSON(http.StatusOK, gin.H{
		"body": "blog deleted successfully with id: " + strconv.Itoa(id),
	})
}

func (gs ginControllerServer) saveBlog(ctx *gin.Context) {
	var requestBody blog

	// Bind the request body to the getRequestBody struct
	if err := ctx.Bind(&requestBody); err != nil {
		// Handle error, e.g., return an error response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := gs.dataStorage.post(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": "succesfull save the blog",
	})
}
