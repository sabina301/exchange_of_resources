package main

import (
	"flag"
	"github.com/sabina301/exchange_of_resources/subjects_manager/db"
	"github.com/sabina301/exchange_of_resources/subjects_manager/rest"
)

func main() {
	port := flag.String("port", "8002", "port to serve on")
	flag.Parse()

	db.InitDB()

	server := rest.NewServer(*port)
	server.Start()
}
