package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":1111", "cert.pem", "key.pem", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi,this is an example of http service in golang!")
}
