// package main
//
// import (
//
//	"bufio"
//	"encoding/json"
//	"fmt"
//	openai "github.com/go-resty/resty/v2"
//	"log"
//	"os"
//
// )
//
//	type Prompt struct {
//		Models      string  `json:"models"`
//		Prompt      string  `json:"prompt"`
//		MaxTokens   int     `json:"max_tokens"`
//		N           int     `json:"n"`
//		Stop        string  `json:"stop"`
//		Temperature float64 `json:"temperature"`
//	}
//
//	type Response struct {
//		Choices []struct {
//			Text string `json:"text"`
//		} `json:"choices"`
//	}
//
//	func main() {
//		// OpenAI APIの認証キーを設定します
//		apiKey := "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"
//		client := openai.New()
//
//		// 翻訳のループを開始します
//		scanner := bufio.NewScanner(os.Stdin)
//		fmt.Print("質問を入力してください（終了するには 'q' を入力）: ")
//		for scanner.Scan() {
//			input := scanner.Text()
//			if input == "q" {
//				break
//			}
//
//			// // APIリクエストの作成
//			// request := client.R().
//			// 	SetHeader("Authorization", "Bearer "+apiKey).
//			// 	SetFormData(map[string]string{
//			// 		"text":        input,
//			// 		"target_lang": "ja",
//			// 	})
//			//
//			// // 翻訳APIの呼び出し
//			// // response, err := request.Post("https://api.openai.com/v1/translations")
//			// response, err := request.Post("https://api.openai.com/v1/engines/davinci-codex/completions")
//			// if err != nil {
//			// 	fmt.Println("Translation error:", err)
//			// 	return
//			// }
//			// fmt.Println("翻訳結果:", response.String())
//
//			prompt := Prompt{
//				Models:      "text-davinci-003",
//				Prompt:      "Translate the following English text to Japanese: 'Hello, how are you?'",
//				MaxTokens:   7,
//				N:           1,
//				Stop:        "\n",
//				Temperature: 0,
//			}
//			jsonData, err := json.Marshal(prompt)
//			if err != nil {
//				fmt.Println("Error marshalling data:", err)
//				return
//			}
//
//			resp, err := client.R().
//				SetHeader("Content-Type", "application/json").
//				SetHeader("Authorization", "Bearer "+apiKey).
//				SetBody(jsonData).
//				Post("https://api.openai.com/v1/completions")
//			// https: // api.openai.com/v1/engines/davinci-codex/completions
//
//			if err != nil {
//				fmt.Println("Error making API request:", err)
//				return
//			}
//
//			if resp.StatusCode() != 200 {
//				fmt.Println("API request failed with status code:", resp.StatusCode())
//				return
//			}
//
//			var response Response
//			err = json.Unmarshal(resp.Body(), &response)
//			if err != nil {
//				fmt.Println("Error unmarshalling response:", err)
//				return
//			}
//
//			if len(response.Choices) > 0 {
//				translatedText := response.Choices[0].Text
//				fmt.Println("Translated text:", translatedText)
//			} else {
//				fmt.Println("No translation found")
//			}
//
//			fmt.Print("質問を入力してください（終了するには 'q' を入力）: ")
//		}
//
//		if err := scanner.Err(); err != nil {
//			log.Fatal("標準入力読み取りエラー:", err)
//		}
//	}
//
// package main
//
// import (
//
//	"context"
//	"fmt"
//	"google.golang.org/api/translate/v3"
//
// )
//
//	func main() {
//		// 翻訳したいテキストを取得します。
//		text := "Hello, world!"
//
//		// クライアントを作成します。
//		client, err := translate.NewClient(context.Background())
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		// リクエストを作成します。
//		request := translate.TranslateTextRequest{
//			// 翻訳元の言語を指定します。
//			SourceLanguageCode: "en",
//			// 翻訳先の言語を指定します。
//			TargetLanguageCode: "fr",
//			// 翻訳するテキストを指定します。
//			Text: text,
//		}
//
//		// 翻訳を取得します。
//		response, err := client.TranslateText(context.Background(), &request)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		// 翻訳されたテキストを出力します。
//		fmt.Println(response.TranslatedText)
//	}
package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/translate/v2"
)

func main() {
	// Google翻訳APIクライアントのセットアップ
	client, err := translate.NewService(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// APIリクエストの作成
	request := &translate.TranslateTextRequest{
		Q:      []string{"Hello, how are you?"},
		Source: "en",
		Target: "fr",
	}

	// APIリクエストの送信
	// response, err := client.Translations.TranslateText(request).Do()
	response, err := client.Translations.Translate(request).Do()

	if err != nil {
		log.Fatal(err)
	}

	// 翻訳結果の表示
	if len(response.Translations) > 0 {
		fmt.Println("Translation:", response.Translations[0].TranslatedText)
	} else {
		fmt.Println("No translation found.")
	}
}
