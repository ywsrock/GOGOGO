package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Note that you need to set the corresponding binding tag on all fields you want to bind.
// For example when binding from JSON, set json:"fieldname".

// Also, Gin provides two sets of methods for binding:
// Type - Must bind
// Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML, BindHeader
// Behavior - These methods use MustBindWith under the hood.
// If there is a binding error, the request is aborted with
// c.AbortWithError(400, err).SetType(ErrorTypeBind).
// This sets the response status code to 400 and the Content-Type header is set to text/plain; charset=utf-8. Note that if you try to set the response code after this, it will result in a warning [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422. If you wish to have greater control over the behavior, consider using the ShouldBind equivalent method.

// Type - Should bind
// Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML, ShouldBindHeader
// Behavior - These methods use ShouldBindWith under the hood. If there is a binding error,
// the error is returned and it is the developer's responsibility to handle the request and error appropriately.

// Login strcut
type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	//router 作成
	r := gin.Default()

	//JSON から　struct
	//curl -v -X POST -H "content-type:application/json" -d {"user":"user","password":"passwd"} http://localhost:8080/loginJson
	r.POST("/loginJson", func(c *gin.Context) {
		var json Login
		//Binding処理
		err := c.ShouldBind(&json)
		//エラー処理
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//ユーザ認証
		if json.User != "user" || json.Password != "passwd" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	//XML からstruct
	r.POST("/loginXml", func(c *gin.Context) {
		var xml Login

		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "user" || xml.Password != "passwd" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, "XML binding")

	})

	r.POST("/loginForm", func(c *gin.Context) {
		var form Login

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "user" || form.Password != "passwd" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "login OK"})
	})

	r.Run()
}
