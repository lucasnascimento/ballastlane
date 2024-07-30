package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lucasnascimento/ballastlane/pkg/clock"
	"github.com/lucasnascimento/ballastlane/pkg/config"
	"github.com/lucasnascimento/ballastlane/pkg/db"
	"github.com/lucasnascimento/ballastlane/pkg/repository"
	"github.com/lucasnascimento/ballastlane/pkg/tools"
)

func main() {

	// Read the configuration path from the command line arguments
	configPath, err := tools.ReadCmdParameters(os.Args)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Create a new configuration manager
	cm, err := config.NewConfigManager(configPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Start the HTTP server to update the configuration
	go func() {
		http.HandleFunc("/update-config", cm.UpdateConfigHandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// Connect to the database
	database, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer db.CloseDB(database)

	// Create a new signals store
	signalStore := repository.NewSignalStore(database)

	// Create a new clock instance
	clock, err := clock.NewClock(cm, clock.RealTimeProvider{}, clock.RealExitProvider{}, signalStore)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Start the clock
	go clock.Run()

	// Waits for an interrupt signal (Ctrl+C) to terminate the program
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")
}
