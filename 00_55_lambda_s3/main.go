package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
)

// awsのリージョン名
var awsRegion = "ap-northeast-1"

type Request struct {
	Name any `json:"BodyName"`
	Age  any `json:"BodyAge"`
}

type Response struct {
	ResName  string
	ResAge   string
	Contents string
}

// S3 内容読み込み
func GetS3Contents(cfg aws.Config) (string, error) {
	s3Info := s3.GetObjectInput{
		// s3 bucket Name
		Bucket: aws.String("s3image-test"),
		// object name
		Key: aws.String("init.txt"),
	}
	// s3 client
	client := s3.NewFromConfig(cfg)
	out, err := client.GetObject(context.TODO(), &s3Info)
	if err != nil {
		return "", err
	}
	body := out.Body
	defer body.Close()
	// read file
	b, err := io.ReadAll(body)
	log.Printf("content -- %s", string(b))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(b), nil
}

func Handler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (Response, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))
	fmt.Printf("Body = %s.\n", request.Body)

	req := Request{}
	if len(request.Body) != 0 {
		// Http Post
		json.Unmarshal([]byte(request.Body), &req)
	} else {
		// Http Get
		req.Name = request.QueryStringParameters["Name"]
		req.Age = request.QueryStringParameters["Age"]
	}

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	fmt.Printf("%#v", request)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	content, err := GetS3Contents(cfg)
	if err != nil {
		log.Fatal(err)
	}
	res := Response{
		ResName:  fmt.Sprintf("%s", req.Name),
		ResAge:   fmt.Sprintf("%s", req.Age),
		Contents: content,
	}
	return res, nil
}
func main() {
	lambda.Start(Handler)
}
