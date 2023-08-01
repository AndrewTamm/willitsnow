package main

import (
	"github.com/AndrewTamm/WillItSnow/cmd"
	"github.com/AndrewTamm/WillItSnow/config"
	"log"
)

func main() {
	cfgPath, err := config.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	server := cmd.NewServer(cfg)

	err = server.Serve()
	if err != nil {
		return
	}
}
