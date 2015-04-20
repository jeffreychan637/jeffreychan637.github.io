package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(r.URL.Path)
	http.ServeFile(w, r, "index.html")

}

func client(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	// http.HandleFunc("/client/", client)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}