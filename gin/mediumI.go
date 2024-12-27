package main

import "github.com/gin-gonic/gin"

type medium interface {
	getBlogs(ctx *gin.Context)
	removeBlog(ctx *gin.Context)
	editBlog(ctx *gin.Context)
	saveBlog(ctx *gin.Context)
}

type dataLayer interface {
	get(int) blog
	post(blog) blog
	edit(blog) blog
	delete(int) interface{}
}
