package main

import (
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

	var files []string
	root := "./"
	grayroot := "./gray"

	//read dir
	filepath.Walk(root,lsFiles(&files))
	for _,file := range files {
		// image to text
		if strings.ToLower(filepath.Ext(file)) == ".png" {
			grayscale(file,grayroot)
		}
	}
}

/**
get contains from image
*/
func img2Text(imgPath string) {
	//ocr client
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("jpn","eng","chi_sim")
	client.SetImage(imgPath)
	//get text
	text, err := client.Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

/**
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

func  grayscale(imgpath string,root string) {
	ff ,err := os.Open(imgpath)
	if err !=nil {
		log.Fatal(err)
	}
	defer ff.Close()
	img,_ := png.Decode(ff)
	bounds := img.Bounds()
	dest := image.NewGray(bounds)

	for y:= bounds.Min.Y;y<=bounds.Max.Y;y++ {
		for x:=bounds.Min.X;x<=bounds.Max.X;x++ {
			c := color.GrayModel.Convert(img.At(x,y))
			gray,_:= c.(color.Gray)
			dest.SetGray(x,y,gray)
		}
	}
	info,_:= os.Stat(imgpath)
	imgName := info.Name()
	f,err := os.Create(fmt.Sprintf("%s/%s",root,imgName))
	if  err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png.Encode(f,dest)

	img2Text(filepath.Join(root, imgName))
}
