package main

import (
	"fmt"
	"log"
	"os"

	"ex.ex/ex/internal/app/apiserver"
	"ex.ex/ex/internal/config"
)

func main() {
	config := config.NewConfig()

	if os.Getenv("PORT") != "" {
		config.BindAddr = ":" + os.Getenv("PORT")
	}

	if os.Getenv("ENDPOINT") != "" {
		config.EndPoint = os.Getenv("ENDPOINT")
	}

	if os.Getenv("DBURL") != "" {
		config.DBURL = os.Getenv("DBURL")
	}

	if os.Getenv("DBNAME") != "" {
		config.DBName = os.Getenv("DBNAME")
	}

	if os.Getenv("COLLECTION") != "" {
		config.Collection = os.Getenv("COLLECTION")
	}

	fmt.Printf("Listen %s\n", config.BindAddr)
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
