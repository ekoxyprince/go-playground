package main

import (
	"fmt"
	"net/http"

	"astro.dev/nethttp-files/handlers"
)

func main() {
	http.HandleFunc("/", handlers.GetHome)
	http.HandleFunc("/upload", handlers.PostUpload)
	fmt.Println("Server is running on PORT 8083")
	http.ListenAndServe(":8083", nil)
}
