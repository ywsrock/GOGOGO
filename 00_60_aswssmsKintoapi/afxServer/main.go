package main

import (
	"afxServer/handler"
	"net/http"
	"time"

	"context"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var Routes = []Route{
	{Method: "GET", Path: "/", Handler: handler.Index},
	//新規作成
	{Method: "POST", Path: "/", Handler: handler.NewRequest},
	//1件取得
	{Method: "GET", Path: "/history", Handler: handler.GetHistory},
	//全レコード
	{Method: "GET", Path: "/historys", Handler: handler.ShowHistory},
	//指定IDで更新
	{Method: "PUT", Path: "/updatehis", Handler: handler.UpdateRequest},
}

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	r.Use(cors.Default())

	for _, h := range Routes {
		r.Handle(h.Method, h.Path, h.Handler)
	}

	ginLambda = ginadapter.New(r)
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// lambda.Start(HandleRequest)

	r := gin.Default()
	r.Use(cors.Default())

	for _, h := range Routes {
		r.Handle(h.Method, h.Path, h.Handler)
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
