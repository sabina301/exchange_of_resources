package main

import (
	"flag"
	"github.com/sabina301/exchange_of_resources/resources_manager/db"
	"github.com/sabina301/exchange_of_resources/resources_manager/repo"
	"github.com/sabina301/exchange_of_resources/resources_manager/rest"
	"os"
)

func main() {
	port := flag.String("port", os.Getenv("PORT"), "port to serve on")
	flag.Parse()
	dbInstance := db.InitDB()
	repoImpl := repo.NewResourceRepository(dbInstance)
	resController := rest.NewResourceController(repoImpl)
	server := rest.NewServer(*port, repoImpl, resController)
	server.Start()
}
