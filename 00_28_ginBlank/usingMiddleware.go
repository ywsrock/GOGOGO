package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Blank gin without middleware")
	//blank gin instance
	r := gin.New()
	//global logger middleware,write log to os.Stdout
	r.Use(gin.Logger())
	//Recovery middleware recovers from any panic
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "Blank Gin")
	})

	r.Run()
}
