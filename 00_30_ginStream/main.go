package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func main() {
	html := `<h1>234</h1>`
	r := gin.Default()
	ctx, timeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeout()

	r.GET("/", hand(ctx, html))
	r.Run(":8989")

}

func hand(ctx context.Context, str string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Stream(func(w io.Writer) bool {
			select {
			case <-ctx.Done():
				fmt.Println("timeout>", ctx.Err())
				return false
			default:
				fmt.Fprint(w, str)
				return true
			}
		})
	}
}
