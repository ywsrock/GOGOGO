# GO_GinとAWS LambdaとAPI Gateway



AWS LambdaとAPI Gatewayを使用してGinフレームワークを実行するためのAWS Lambda関数です。AWS LambdaとGinフレームワークを組み合わせることで、サーバーレスな環境で高性能なAPIを構築することができます。



> `github.com/awslabs/aws-lambda-go-api-proxy/gin` パッケージは、AWS LambdaとGinフレームワークを統合するためのアダプターです。
>
> AWS Lambdaで実行される関数をGinフレームワークとして扱うことができます。
>
> 以下の機能を提供します：
>
> - `New` 関数：Ginフレームワークのルーターを受け取り、Lambdaのハンドラー関数として実行するための `http.HandlerFunc` を返します。
> - `NewWithContext` 関数：`New` 関数と同様の機能を提供しますが、追加で `context.Context` を受け取ることができます。
> - `NewGinLambda` 関数：Ginフレームワークのルーターを受け取り、AWS Lambdaで実行するためのハンドラー関数を返します。
>
> これらの関数を使用することで、Ginフレームワークで定義されたエンドポイントやミドルウェアをAWS Lambda関数として実行することができます。AWS API GatewayなどのイベントソースからのリクエストがLambda関数に渡され、Ginフレームワークのルーターでリクエストが処理されます。



## コードの概要

コードの詳細を説明します。



### パッケージのインポート

```go
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)
```

このセクションでは、必要なパッケージをインポートしています。

- `github.com/aws/aws-lambda-go/events` パッケージは、AWS Lambda関数のイベントとレスポンスの構造体を提供します。
- `github.com/aws/aws-lambda-go/lambda` パッケージは、AWS Lambda関数のハンドラーを定義するための関数を提供します。
- `github.com/awslabs/aws-lambda-go-api-proxy/gin` パッケージは、GinフレームワークをAWS LambdaとAPI Gatewayに統合するためのアダプターを提供します。
- `github.com/gin-gonic/gin` パッケージは、Ginフレームワーク自体の機能を提供します。

### RouteとRoutes

```go
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
```

`Route` 構造体は、HTTPメソッド、パス、およびハンドラー関数を表します。`Routes` は `Route` 構造体のスライスであり、APIエンドポイントの定義を格納しています。

### init関数

```go
var ginLambda *ginadapter.GinLambda

func init() {
	log.Printf("Gin cold start")
	r := gin.Default()

	for _, h := range Routes {
		r.Handle(h.Method, h.Path, h.Handler)
	}

	ginLambda = ginadapter.New(r)
}
```

`init` 関数は、Lambda関数の初期化時に実行されます。Ginフレームワークの初期化、APIエンドポイントのルーティング設定、およびGinLambdaアダプターの作成が行われます。このアダプターは、GinフレームワークをAWS LambdaとAPI Gatewayに統合するためのものです。

### f関数

```go
func f(c *gin.Context) {
	c.JSON(200, gin.H{
		

	"message": "pong",
		})
}
```

`f` 関数は、Ginフレームワークのハンドラー関数です。この例では、HTTPリクエストに対して "pong" メッセージを含むJSONレスポンスを返します。

### Handler関数

```go
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}
```

`Handler` 関数は、AWS Lambdaのハンドラー関数です。この関数は、GinLambdaアダプターを使用してGinフレームワークをAWS LambdaとAPI Gatewayに統合します。

### main関数

```go
func main() {
	lambda.Start(Handler)
}
```



`main` 関数は、アプリケーションのエントリーポイントです。AWS Lambda関数の開始と、`Handler` 関数の実行が行われます。