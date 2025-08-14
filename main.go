package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/ygzaydn/go-dnswatcher/internal/config"
	"github.com/ygzaydn/go-dnswatcher/internal/dnsHandler"
	"github.com/ygzaydn/go-dnswatcher/internal/gui"
)

func main() {
	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
		os.Exit(1)
	}

	guiMode := flag.Bool("gui", false, "Run in GUI mode")
	flag.Parse()

	if *guiMode {
		log.Println("Starting in GUI mode...")
		log.SetOutput(io.Discard)
		go dnsHandler.Start(*cfg)
		gui.Start()
	} else {
		log.Println("Starting in CLI mode...")
		dnsHandler.Start(*cfg)
	}
	log.Println("DNSWatcher is running. Press Ctrl+C to exit.")
}
