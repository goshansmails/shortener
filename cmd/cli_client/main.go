package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goshansmails/shortener/internal/client"
)

const defaultAddress = "http://127.0.0.1:8080"

const (
	actionShorten = "s"
	actionLonger  = "l"
)

var (
	action         = flag.String("action", actionLonger, "action: 's' for 'shorten'; 'l' for 'longer'")
	serverHostPort = flag.String("address", defaultAddress, "address of shortener server")
	url            = flag.String("url", "", "url to process")
)

func main() {

	flag.Parse()

	client := client.New(*serverHostPort)

	var actFunc func(string) (string, error)

	switch *action {
	case actionShorten:
		actFunc = client.ShortenURL
	case actionLonger:
		actFunc = client.LongerURL
	default:
		fmt.Println("bad action:", *action)
		os.Exit(1)
	}

	result, err := actFunc(*url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}
