package handler

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddBlogAPI(c *gin.Context) {
	title := c.Request.FormValue("title")
	author := c.Request.FormValue("author")
	article := c.Request.FormValue("article")

	blog := models.Blog{Title: title, Author: author, Article: article}
	ra_rows, err := blog.AddBlog()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Insert successful! %d", ra_rows)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetBlogsAPI(c *gin.Context) {
	var b models.Blog
	blogs, err := b.GetBlogs()

	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"blogs": blogs,
	})
}
func GetBlogAPI(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	b := models.Blog{Id: id}
	blogs, err := b.GetBlogs()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"blogs": blogs,
	})
}

func UpdateBlogAPI(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	b := models.Blog{Id: id}
	err = c.Bind(&b)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := b.UpdateBlog()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("Update successful! %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeleteBlogAPI(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	b := models.Blog{Id: id}
	ra, err := b.DeleteBlog()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful %d!", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
