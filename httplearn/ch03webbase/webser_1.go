package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method", r.Method)
	fmt.Println("host : ", r.Host)
	fmt.Println("url : ", r.URL)
	fmt.Println("url.Path", r.URL.Path)
	fmt.Println("url.Scheme", r.URL.Scheme)
	fmt.Println("header : ", r.Header)
	fmt.Println("form : ", r.Form)
	for k, v := range r.Form {
		fmt.Println(k, " : ", v)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9090", nil)
}
