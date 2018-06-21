package main

import (
	"net/http"
	"io"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

func main() {
		http.HandleFunc("/", HelloServer)
		log.Fatal(http.ListenAndServe(":8100", nil))
}