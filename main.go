package main

import (
	"QRMenu/config"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	config.InitDB()

	http.HandleFunc("/menu", handler)
	http.ListenAndServe(":8080", nil)
}
