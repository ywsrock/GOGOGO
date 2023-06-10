package awssms

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

const (
	AwsAccessKey       = "your key"
	AwsSecretAccessKey = "your key"
	AwsRegion          = "your region"
)

type SnsClient struct{}

func PublishMessage(svc snsiface.SNSAPI, input *sns.PublishInput) (*sns.PublishOutput, error) {
	result, err := svc.Publish(input)
	return result, err
}

func (c *SnsClient) SendSms(msg string, phoneNo string) (string, error) {
	// クライアントの生成
	client, err := c.NewSNSClient()
	if err != nil {
		return "", nil
	}

	// メッセージの作成
	pin := &sns.PublishInput{}
	pin.SetMessage(msg)
	pin.SetPhoneNumber(phoneNo)

	result, err := PublishMessage(client, pin)
	if err != nil {
		fmt.Printf("%s%s\n", "SMS送信エラー:", err)
		return "", err
	}

	return *result.MessageId, nil
}

func (c *SnsClient) NewSNSClient() (*sns.SNS, error) {
	creds := credentials.NewStaticCredentials(AwsAccessKey, AwsSecretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(AwsRegion),
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	client := sns.New(sess)
	return client, nil
}
