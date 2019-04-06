package main

import (
	"log"
	"net/http"

	"example/sample"
)

func main() {
	http.HandleFunc("/", sample.Handler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
