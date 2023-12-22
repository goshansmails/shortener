package main

import (
	"fmt"
	"net/http"
	"os"
)

func empty(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, empty)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		fmt.Printf("cat't run server: %s\n", err.Error())
		os.Exit(1)
	}
}
