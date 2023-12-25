package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goshansmails/shortener/internal/app/server"
	"github.com/goshansmails/shortener/internal/app/store/mapstore"
)

var (
	addr    = flag.String("a", "0.0.0.0:8080", "address to listen")
	baseURL = flag.String("b", "http://127.0.0.1:8080", "base URL for redirection")
)

func main() {

	flag.Parse()

	settings := server.Settings{
		Addr:    *addr,
		BaseURL: *baseURL,
		Store:   mapstore.New(),
	}

	s := server.New(settings)
	if err := s.Run(); err != nil {
		fmt.Println("can't run server: ", err)
		os.Exit(1)
	}
}
