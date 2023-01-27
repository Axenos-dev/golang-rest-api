package main

import (
	"log"
	"mazano-server/src/internal/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()

	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
