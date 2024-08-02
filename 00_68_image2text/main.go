package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
)

// 画像の前処理を行う関数
func preprocessImage(inputPath string, outputPath string) error {
	imgFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	// 最小限の前処理
	img = imaging.AdjustBrightness(img, 10)
	img = imaging.AdjustContrast(img, 10)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// JPEG 形式で保存
	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		return err
	}

	return nil
}

// テキストをファイルに保存する関数
func saveTextToFile(text string, outputPath string) error {
	err := ioutil.WriteFile(outputPath, []byte(text), 0o644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	preFolder := "image/pre"
	postFolder := "image/post"
	textFolder := "text"

	if err := os.MkdirAll(postFolder, os.ModePerm); err != nil {
		fmt.Println("Postフォルダー作成エラー:", err)
		return
	}
	if err := os.MkdirAll(textFolder, os.ModePerm); err != nil {
		fmt.Println("Textフォルダー作成エラー:", err)
		return
	}

	files, err := ioutil.ReadDir(preFolder)
	if err != nil {
		fmt.Println("Preフォルダー読み込みエラー:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		inputPath := filepath.Join(preFolder, file.Name())
		postPath := filepath.Join(postFolder, file.Name())
		fileNameWithoutExt := file.Name()[:len(file.Name())-len(filepath.Ext(file.Name()))]
		textFilePath := filepath.Join(textFolder, fileNameWithoutExt+".txt")

		// 画像前処理
		err := preprocessImage(inputPath, postPath)
		if err != nil {
			fmt.Println("前処理エラー:", err)
			continue
		}

		// 前処理後の画像を手動で確認
		fmt.Println("前処理後の画像保存先:", postPath)

		// Tesseract OCR クライアントの作成
		client := gosseract.NewClient()
		defer client.Close()

		client.SetLanguage("jpn+eng+chi_sim+chi_tra")
		client.SetImage(postPath)
		// PSMの設定
		client.SetPageSegMode(gosseract.PSM_AUTO_OSD) // PSM: 自動ページセグメンテーション

		// テキスト抽出
		text, err := client.Text()
		if err != nil {
			fmt.Println("OCR エラー:", err)
			continue
		}

		// デバッグ用にテキストをコンソールに出力
		fmt.Println("抽出したテキスト:", text)

		// テキストをファイルに保存
		if err := saveTextToFile(text, textFilePath); err != nil {
			fmt.Println("テキストファイル保存エラー:", err)
			continue
		}

		fmt.Println("抽出したテキストを保存しました:", textFilePath)
	}
}
