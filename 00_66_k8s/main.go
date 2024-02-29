package main

import "net/http"
import "os"

func handler(w http.ResponseWriter, r *http.Request) {
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "not exist"
	}
	w.Write([]byte(greeting + "Hello, World!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
