package main

import (
	"database/sql"
	"log"
	"net/http"

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

	type Blgo struct {
		Id      int    `json:"id" form:"id"`
		Title   string `json:"title" form:"title"`
		Author  string `json:"author" form:"author"`
		Article string `json:"article" form:"article"`
	}

	router.GET("/blgos", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM blogarticle")

		if err != nil {
			log.Fatalln(err)
		}
		defer rows.Close()

		blgos := make([]Blgo, 0)

		for rows.Next() {
			var blgo Blgo
			rows.Scan(&blgo.Id, &blgo.Title, &blgo.Author, &blgo.Article)
			blgos = append(blgos, blgo)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"blgos": blgos,
		})

	})

	router.Run(":8000")

}
