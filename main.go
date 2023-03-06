package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", listener)
	fmt.Println("Starting....")
	http.ListenAndServe(":8080", nil)
}
