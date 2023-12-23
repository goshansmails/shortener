package main

import (
	"fmt"
	"os"

	"github.com/goshansmails/shortener/internal/app/server"
	"github.com/goshansmails/shortener/internal/app/store/mapstore"
)

func main() {

	s := server.New(mapstore.New())
	if err := s.Run(); err != nil {
		fmt.Println("can't run server: ", err)
		os.Exit(1)
	}
}
