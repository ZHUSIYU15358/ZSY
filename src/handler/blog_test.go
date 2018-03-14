package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	//初始化路由
	router := gin.Default()

	defer database.SQLdb.Close()

	router.GET("/getblogs", GetBlogsAPI)

	router.GET("/getblog", GetBlogAPI)

	router.GET("/addblog", AddBlogAPI)

	router.GET("/updateblog", UpdateBlogAPI)

	router.GET("/deleteblog", DeleteBlogAPI)

	router.Run(":8000")

}

func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}

	temp := values[1:]
	values = "?" + temp
	return values
}

func JsonRequest(method string, uri string, param interface{},
	router *gin.Engine) []byte {

	jsonByte, _ := json.Marshal(param)

	req := httptest.NewRequest(method, uri,
		bytes.NewReader(jsonByte))

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	body, _ := ioutil.ReadAll(result.Body)
	return body
}

type Req struct {
	message string `json:"msg"`
}

func TestAddBlogAPI(t *testing.T) {
	url := "/addblog"
	param := make(map[string]string)
	param["Title"] = "test123"
	param["Author"] = "zsyyy"
	param["Article"] = "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."
	body := FormRequest("GET", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	str := "Insert successful!"
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != str {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func TestGetBlogsAPI(t *testing.T) {
	url := "/getblogs"
	param := make(map[string]string)
	body := FormRequest("GET", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != "success" {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func TestGetBlogAPI(t *testing.T) {
	url := "/getblog"
	param := make(map[string]string)
	param["Id"] = "3"
	body := FormRequest("GET", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != "success" {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}
func TestUpdateBlogAPI(t *testing.T) {
	url := "/updateblog"
	param := make(map[string]string)
	param["Id"] = "6"
	body := FormRequest("GET", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	str := "Update successful!"
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != str {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}
func TestDeleteBlogAPI(t *testing.T) {
	url := "/deleteblog"
	param := make(map[string]string)
	param["Id"] = "6"
	body := FormRequest("GET", url, param, router)
	fmt.Printf("response:%v\n", string(body))
	str := "Delete successful!"
	r := &Req{}
	err := json.Unmarshal(body, r)
	if r.message != str {
		t.Errorf(" 响应字符串不符，body:%v\n", string(body))
	}
	if err != nil {
		log.Fatalln(err)
	}
}
