#  Go_AWSCDK

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests





## コードの概要

AWS CDKを使用してAWSのリソースをプロビジョニングするためのアプリケーションです。具体的には、AWS Lambda関数とAPI Gatewayを作成しています。

### パッケージのインポート

```go
package main

import (
    "github.com/aws/aws-cdk-go/awscdk/v2"
    "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
    "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
    "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
)
```

まず、使用するパッケージをインポートしています。以下は各パッケージの役割です：

- `github.com/aws/aws-cdk-go/awscdk/v2`：AWS CDKのコアライブラリです。AWSリソースをプロビジョニングするためのクラスや関数が提供されています。
- `github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2`：AWS CDKのLambda Go Runtime用のライブラリです。Go言語で記述されたAWS Lambda関数を作成するためのクラスや関数が提供されています。
- `github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway`：AWS CDKのAPI Gateway用のライブラリです。API Gatewayの設定やリソースを作成するためのクラスや関数が提供されています。
- `github.com/aws/constructs-go/constructs/v10`：ConstructsフレームワークのGo版です。AWS CDKで使用されるコンストラクトの基礎となるクラスや関数が提供されています。
- `github.com/aws/jsii-runtime-go`：AWS CDKのJavaScriptインターフェースインフラストラクチャ（JSII）ランタイムのGo版です。CDKアプリケーションとAWSの間での相互作用をサポートしています。



### スタックの定義

```go
type HelloStackProps struct {
    awscdk.StackProps
}

func NewCfdStack(scope constructs.Construct, id string, props *HelloStackProps) awscdk.Stack {
    // スタックの初期化
    // ...

    // Lambda関数の作成
    // ...

    // API Gatewayの作成
    // ...

    return stack
}
```

`HelloStackProps` 構造体は、CDKのスタックのプロパティを定義するためのものです。`NewCfdStack` 関数は、CDKのスタックを作成するための関数です。このスタックでは、AWS Lambda関数とAPI Gatewayが作成されます。



### Lambda関数の作成

```go
func NewAfLambdaHandler(stack awscdk.Stack) awscdklambdagoalpha.GoFunction {
    // Lambda関数の初期化
    // ...

    return handler
}
```

`NewAfLambdaHandler` 関数は、AWS Lambda関数を作成するための関数です。この関数は、指定されたディレクトリからGo言語で書かれたコードを読み込み、AWS Lambda関数としてデプロイします。（※LambdaとGin統合利用しています。Lambdaフォルダーを参照）



### API Gatewayの作成

```go
func NewAfApiGatewap(stack awscdk.Stack, handler awscdklambdagoalpha.GoFunction) {
    // API Gatewayの初期化
    // ...

    // API Gatewayのリソースとメソッドの追加
    // ...
}
```

`NewAfApiGatewap` 関数は、AWS API Gatewayを作成するための関数です。API GatewayはRESTfulなエンドポイントを提供し、Lambda関数と統合します。この関数内では、API Gatewayの初期化やエンドポイントの設定が行われます。



### メイン関数

```go
func main() {
    // アプリケーションの初期化
    // ...

    // スタックの作成
    // ...

    // デプロイ
    // ...
}
```

`main` 関数は、CDKアプリケーションのエントリーポイントです。CDKアプリケーションを初期化し、スタックを作成します。CDKアプリケーションを構築した後、`app.Synth()` メソッドを呼び出すことで、AWS CloudFormationテンプレートが生成されます。



### 環境設定

```go
func env() *awscdk.Environment {
    // 環境の設定
    // ...

    return nil
}
```

`env` 関数は、CDKスタックをデプロイするAWSの環境（アカウントとリージョン）を設定するための関数です。現在のCLIの設定や環境変数を基に環境を指定することができます。