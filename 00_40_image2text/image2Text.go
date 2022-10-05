package main

import (
	"flag"
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var files []string

func main() {
	// var files []string
	var root string
	defPath, _ := os.Getwd()
	flag.StringVar(&root, "image", defPath, "Parse Image Path")
	flag.Parse()

	grayRoot := filepath.Join(defPath, "gray")

	// read dir
	err := filepath.Walk(root, lsFiles(&files))
	if err != nil {
		return
	}
	for _, file := range files {
		// image to text
		if strings.ToLower(filepath.Ext(file)) == ".png" {
			grayscale(file, grayRoot)
		}
	}
	err = os.RemoveAll(grayRoot)
	if err != nil {
		return
	}
}

/*
*
get contains from image
*/
func img2Text(imgPath string) {
	// ocr client
	client := gosseract.NewClient()
	defer func(client *gosseract.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	// 識別言語
	err := client.SetLanguage("jpn", "eng", "chi_sim")
	if err != nil {
		return
	}
	err = client.SetImage(imgPath)
	if err != nil {
		return
	}
	// get text
	text, err := client.Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

/*
*
再帰的にファイルを検索
*/
func lsFiles(files *[]string) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func grayscale(imgPath string, root string) {
	ff, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(ff *os.File) {
		err := ff.Close()
		if err != nil {
			return
		}
	}(ff)

	if _, err := os.Stat(root); os.IsNotExist(err) {
		err := os.Mkdir(root, 0777)
		if err != nil {
			return
		}
	}

	img, _ := png.Decode(ff)
	bounds := img.Bounds()
	dest := image.NewGray(bounds)

	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y))
			gray, _ := c.(color.Gray)
			dest.SetGray(x, y, gray)
		}
	}

	info, _ := os.Stat(imgPath)
	imgName := info.Name()
	f, err := os.Create(fmt.Sprintf("%s/%s", root, imgName))
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	err = png.Encode(f, dest)
	if err != nil {
		return
	}
	img2Text(filepath.Join(root, imgName))
}
