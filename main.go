package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	type Blog struct {
		Id      int    `json:"id" form:"id"`
		Title   string `json:"title" form:"title"`
		Author  string `json:"author" form:"author"`
		Article string `json:"article" form:"article"`
	}

	router.GET("/getblogs", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM blogarticle")

		if err != nil {
			log.Fatalln(err)
		}
		defer rows.Close()

		blogs := make([]Blog, 0)

		for rows.Next() {
			var blog Blog
			rows.Scan(&blog.Id, &blog.Title, &blog.Author, &blog.Article)
			blogs = append(blogs, blog)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"blogs": blogs,
		})

	})

	router.GET("/getblog/:id", func(c *gin.Context) {
		id := c.Param("id")
		var blog Blog
		err := db.QueryRow("SELECT * FROM blogarticle WHERE id=?", id).Scan(
			&blog.Id, &blog.Title, &blog.Author, &blog.Article,
		)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"blog": nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"blog": blog,
		})

	})

	router.GET("/addblog", func(c *gin.Context) {
		title := c.Request.FormValue("title")
		author := c.Request.FormValue("author")
		article := c.Request.FormValue("article")

		rs, err := db.Exec("INSERT INTO blogarticle(title,author,article) VALUES (?, ?, ?)", title, author, article)
		if err != nil {
			log.Fatalln(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert blogarticle Id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	router.GET("/updateblog/:id", func(c *gin.Context) {
		cid := c.Param("id")
		id, err := strconv.Atoi(cid)
		blog := Blog{Id: id}
		err = c.Bind(&blog)
		if err != nil {
			log.Fatalln(err)
		}

		stmt, err := db.Prepare("UPDATE blogarticle SET title=?, author=?, article=? WHERE id=?")
		defer stmt.Close()
		if err != nil {
			log.Fatalln(err)
		}
		rs, err := stmt.Exec(blog.Title, blog.Author, blog.Article, blog.Id)
		if err != nil {
			log.Fatalln(err)
		}
		ra, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("Update blogarticle %d successful %d", blog.Id, ra)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	router.GET("/deleteblog/:id", func(c *gin.Context) {
		cid := c.Param("id")
		id, err := strconv.Atoi(cid)
		if err != nil {
			log.Fatalln(err)
		}
		rs, err := db.Exec("DELETE FROM blogarticle WHERE id=?", id)
		if err != nil {
			log.Fatalln(err)
		}
		ra, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("Delete blogarticle %d successful %d", id, ra)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	router.Run(":8000")

}
