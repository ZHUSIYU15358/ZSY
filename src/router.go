package main

import (
	"handler"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/getblogs", handler.GetBlogsAPI)

	router.GET("/getblog", handler.GetBlogAPI)

	router.GET("/addblog", handler.AddBlogAPI)

	router.GET("/updateblog", handler.UpdateBlogAPI)

	router.GET("/deleteblog", handler.DeleteBlogAPI)

	return router
}
