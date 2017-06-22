package main

import (
	"os"
	"net/http"
	_ "github.com/katie31/extract"
)

func main() {
	home := os.Getenv("HOME")
	http.HandleFunc("/", extract.Handler)
	//http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS(":8080", home + "/server.crt", home + "/server.key", nil)

	if err != nil {
		panic(err)
	}
}


