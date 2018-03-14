package models

import (
	"database/sql"
	"log"
)

type Blog struct {
	Id      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Author  string `json:"author" form:"author"`
	Article string `json:"article" form:"article"`
}

var db *sql.DB

func (b *Blog) AddBlog() (id int64, err error) {
	rs, err := db.Exec("INSERT INTO blogarticle(title,author,article) VALUES (?, ?, ?)", b.Title, b.Author, b.Article)
	if err != nil {
		return
	}
	id, err1 := rs.LastInsertId()
	if err1 != nil {
		return
	}
	return
}

func (B *Blog) GetBlogs() (blogs []Blog, err error) {
	blogs = make([]Blog, 0)
	rows, err := db.Query("SELECT * FROM blogarticle")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var blog Blog
		rows.Scan(&blog.Id, &blog.Title, &blog.Author, &blog.Article)
		blogs = append(blogs, blog)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (b *Blog) GetBlog() (blog Blog, err error) {
	err2 := db.QueryRow("SELECT * FROM blogarticle WHERE id=?", b.Id).Scan(
		&blog.Id, &blog.Title, &blog.Author, &blog.Article,
	)
	if err2 != nil {
		return
	}
	return
}

func (b *Blog) UpdateBlog() (ra int64, err error) {
	stmt, err := db.Prepare("UPDATE blogarticle SET title=?, author=?, article=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(b.Title, b.Author, b.Article, b.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (b *Blog) DeleteBlog() (ra int64, err error) {
	rs, err := db.Exec("DELETE FROM blogarticle WHERE id=?", b.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()

	return
}
