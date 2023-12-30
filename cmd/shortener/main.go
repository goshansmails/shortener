package main

import (
	"flag"
	"os"

	"github.com/goshansmails/shortener/internal/server"
	"github.com/goshansmails/shortener/internal/store/mapstore"
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
		BaseURL: getBaseURL(),
		Store:   mapstore.New(),
	}

	server.Run(getServerAddress(), settings)
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
