package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var strFilePath string
	strFilePath = "test.txt"
	// readFile by OS.Read
	if err := readFileByBates(strFilePath); err != nil {
		log.Fatal(err)
	}
	// readBy bufio.Scanner
	if err := readFileBybufioScanner(strFilePath); err != nil {
		log.Fatal(err)
	}
	// bufio.ReadLine()
	if err := readFileByBufioReader(strFilePath); err != nil {
		log.Fatal(err)
	}
	// bufio.ReadString(delim)
	if err := readFileByBufioReaderString(strFilePath, ','); err != nil {
		log.Fatal(err)
	}
	// ioutil.ReadAll()
	if err := readFileByIoutilReadAll(strFilePath); err != nil {
		log.Fatal(err)
	}
	//ioutil.ReadFile()
	if err := readFileByIoutilReadFile(strFilePath); err != nil {
		log.Fatal(err)
	}
}

// readfile by os
func readFileByBates(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	bs := make([]byte, 128)
	fmt.Println("os.Read------->")
	for {
		n, err := f.Read(bs)
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(string(bs[:n]))
	}
	return nil
}

// readByBufio.Scanner
func readFileBybufioScanner(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println("bufio.Scanner()---->")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

//bufio.Reader()
func readFileByBufioReader(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	fmt.Println("reader.ReadLine()------>")
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(string(line))
		if !isPrefix {
			fmt.Println()
		}
	}
	return nil
}

//bufio.readerString
func readFileByBufioReaderString(fileName string, delim byte) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println("bufio.ReadString(delim)------->")
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString(delim)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(line)
	}
	return nil
}

//ioutil.ReadAll
func readFileByIoutilReadAll(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println("readAll()----->")
	text, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	fullText := string(text)
	fmt.Println(fullText)
	return nil
}

//iotuil.readFile()
func readFileByIoutilReadFile(fileName string) error {
	fmt.Println("readFile()------>")
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	fmt.Println(string(f))
	return nil
}

// -----write file
