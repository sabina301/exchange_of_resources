package main

import (
	"flag"
	"github.com/sabina301/exchange_of_resources/api_gateway/rest"
)

func main() {
	port := flag.String("port", "8080", "port to serve on")
	flag.Parse()

	server := rest.NewServer(*port)
	server.Start()
}
