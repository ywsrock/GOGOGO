package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(fileName string) *os.File {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	// 長いファイル処理とき、Buffer足りない可能せいがあります
	scanner := bufio.NewScanner(f)
	// 単語単位で読む(default line)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return f
}

func writeFile(fileName string) {
	// f, err := os.Create(fileName)
	f, err := os.OpenFile(fileName, os.O_WRONLY, 666)
	// f.Close()
	if err == nil {
		bw := bufio.NewWriter(f)
		_, err = bw.WriteString("123")
		if err != nil {
			fmt.Println(err)
		}
		// buffer 内容書き込み
		bw.Flush()
		fmt.Println("------999")
	}

	f.Close()
}

func readFromString(str string) {
	s := strings.NewReader(str)
	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := readFile("test.txt")
	fmt.Println(f)
	f.Close()
	var str string
	str = `test
	aasfd
	asdf
	asdf
	asdf
	asdf
	`
	readFromString(str)
	writeFile("out.txt")
}
