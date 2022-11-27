package main

import (
	"flag"
	"log"
	"mazano-server/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

var (
	path_to_config string
)

func init() {
	flag.StringVar(&path_to_config, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(path_to_config, config)

	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
