package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//get処理
	res, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	bs, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bs))
	//POST
	v := url.Values{}
	v.Add("key", "123")
	resP, err1 := http.PostForm("https://www.google.com", v)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(resP.Header)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
