package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"greeting"`
}

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var Routes = []Route{
	{Method: "GET", Path: "/", Handler: f},
	//新規作成
	{Method: "POST", Path: "/", Handler: f},
	//1件取得
	{Method: "GET", Path: "/history", Handler: f},
	//全レコード
	{Method: "GET", Path: "/historys", Handler: f},
	//指定IDで更新
	{Method: "PUT", Path: "/updatehis", Handler: f},
}

var ginLambda *ginadapter.GinLambda

func init() {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()

	for _, h := range Routes {
		r.Handle(h.Method, h.Path, h.Handler)
	}

	ginLambda = ginadapter.New(r)
}

func f(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	// return "Hello, World!", nil
// 	resp := &response{
// 		Message: "hello world!",
// 	}
// 	body, err := json.Marshal(resp)
// 	if err != nil {
// 		return events.APIGatewayProxyResponse{Body: string("Error parsing payload"), StatusCode: 400}, err
// 	}
// 	return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil

// }

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
