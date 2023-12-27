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

func init() {
	flag.Parse()
}

func main() {
	settings := server.Settings{
		Addr:    getServerAddress(),
		BaseURL: getBaseURL(),
		Store:   mapstore.New(),
	}

	s := server.New(settings)
	if err := s.Run(); err != nil {
		fmt.Println("can't run server:", err)
		os.Exit(1)
	}
}

func getServerAddress() string {
	result := os.Getenv("SERVER_ADDRESS")
	if result != "" {
		return result
	}

	return *addr
}

func getBaseURL() string {
	result := os.Getenv("BASE_URL")
	if result != "" {
		return result
	}

	return *baseURL
}
