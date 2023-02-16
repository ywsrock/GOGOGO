package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// r.Static("/", "./public")

	// default
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// path param
	//  curl http://localhost:8080/user/123
	r.GET("/user/:name", func(c *gin.Context) {
		userName := c.Param("name")
		c.String(200, "hello %s", userName)
	})

	// path query
	// curl 'http://localhost:8080/query?id=123&page=666'
	r.GET("/query", func(c *gin.Context) {
		queryID := c.Query("id")
		queryPage := c.Query("page")
		c.String(200, "id=%s page=%s", queryID, queryPage)
	})

	// post Form
	// curl -d message="test" 'http://localhost:8080/post?id=123&page=666'
	r.POST("/post", func(c *gin.Context) {
		message := c.PostForm("message")
		id := c.Query("id")
		page := c.Query("page")
		c.String(200, "POST message=%s id=%s page=%s", message, id, page)
	})

	// ファイルアップロード
	// curl -X POST http://localhost:8080/fileUp    -F "file=@C:\Users\user\Desktop\ang\GO\00_27_ginQuick\aaa\aaa.zip" name="123"  email="a@a.com" -H "Content-Type: multipart/form-data"
	r.POST("/fileUp", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// single file
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		// アップロードファイル名取得
		filename := filepath.Base(file.Filename)
		mdfile := md5.Sum([]byte(filename))
		fmt.Println(file.Filename)

		// 指定パスファイル保存
		pwd, _ := os.Getwd()
		dst := filepath.Join(pwd, string(mdfile[:]))

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
	})

	// Grouping routes
	g1 := r.Group("groupUser/")
	{
		g1.GET("/id", func(c *gin.Context) {
			c.String(200, fmt.Sprintf("user group　%s", c.Query("id")))
		})
	}

	g2 := r.Group("admin/")
	{
		g2.GET("/admin", func(c *gin.Context) {
			c.String(200, fmt.Sprintf("admin group %s", c.Query("admin")))
		})
	}
	r.Run()
}
