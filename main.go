package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===== Remote Address =====")
	fmt.Println(r.RemoteAddr)

	fmt.Println("===== Method =====")
	fmt.Println(r.Method)

	fmt.Println("===== Host =====")
	fmt.Println(r.URL.Host)

	fmt.Println("===== URI =====")
	fmt.Println(r.RequestURI)

	fmt.Println("===== Path =====")
	fmt.Println(r.URL.Path)

	fmt.Println("===== RawQuery =====")
	fmt.Println(r.URL.RawQuery)

	fmt.Println("===== Header =====")
	for k, v := range r.Header {
		fmt.Printf("[%v]:%s\n", k, v)
	}

	fmt.Println("===== Form =====")
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Printf("[%v]:%s\n", k, v)
	}

	fmt.Println("===== Body =====")
	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	fmt.Printf(body.String())
}
