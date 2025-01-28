package main

import (
	"log"
	"os"

	"github.com/ygzaydn/go-dnswatcher/internal/config"
	"github.com/ygzaydn/go-dnswatcher/internal/dnsHandler"
)

func main() {
    cfg, err := config.LoadConfig("config.yaml")
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
        os.Exit(1)
    }

    dnsHandler.Start(*cfg)
}