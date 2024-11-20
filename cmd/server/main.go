package main

import (
	"log"
	"net/http"

	"github.com/leetcode-golang-classroom/golang-multipart-upload-sample/internal/controller"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	addr := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", HelloHandler)
	mux.HandleFunc("POST /upload", controller.FileUpload)
	mux.HandleFunc("POST /upload_multipart", controller.FileUploadMultipart)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
