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
	scanner := bufio.NewScanner(f)
	// 単語単位で読む
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
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
