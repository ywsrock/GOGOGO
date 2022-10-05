package main

import (
	"fmt"
	"net/http"
	"time"
)

func getHandle() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, time.Now().Format("2006/01/02 15:04:05"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getHandle())
	http.ListenAndServe(":9999", mux)
}
