package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
)

/**
　数が多くなると、Routineの数も膨大にんなる
 */


func get (url string) {
	resp,err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Gorution:%d,Url:%s,(%d bytes)",runtime.NumGoroutine(),url,len(body))
}


func main() {
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			url := fmt.Sprintf("https://placehold.jp/%dx%d.png",i,i)
			get(url)
		}(i)
	}
	wg.Wait()
}