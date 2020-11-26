package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

//ログフォマット
func logFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func main() {
	fmt.Println("gin log file")
	//ログカラー無効
	gin.DisableConsoleColor()
	//ログカラー有効
	// gin.ForceConsoleColor()

	//ログファイル作成
	f, err := os.Create("logfile.log")
	if err != nil {
		log.Fatal("ファイル作成エラー")
	}

	//default出力先変更
	gin.DefaultWriter = io.MultiWriter(f)
	//router作成
	r := gin.Default()

	//ログ取得フォマットミドル利用
	r.Use(gin.LoggerWithFormatter(logFormatter))

	//Get 請求
	r.GET("/log", func(c *gin.Context) {
		c.String(200, "write log file")
	})

	r.Run()
}
