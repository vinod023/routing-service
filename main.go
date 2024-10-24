package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside the hello controller")
	fmt.Fprintf(w, "hello go func. It's a router \n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("router successfully started")
	http.ListenAndServe(":8090", nil)
}

// func main() {
// 	fmt.Println("Hello... This go func")
// }
