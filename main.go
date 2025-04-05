package main

import (
	"context"
	"flag"
	"fmt"
	"kvs/core"
	"kvs/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Flags
	portPtr := flag.Int("p", 3000, "unsigned integer port number")
	flag.Parse()

	port := *portPtr
	if port < 1 || port > 65535 {
		fmt.Printf("Invalid port number\n")
		os.Exit(1)
	}
	addr := fmt.Sprintf(":%d", *portPtr)

	// Initialize Server
	server := &http.Server{
		Addr:    addr,
		Handler: handlers.GetMux(),
	}

	core.InitStore()

	// Start Server
	fmt.Printf("\033[1;37mStarting a server on port %d\033[0m\n\n", port)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Could not start the server at port %d\n", port)
			os.Exit(1)
		}
	}()

	// Some Background jobs
	ticker, done := core.StartTicker(30, core.KillExpiredKeys)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan
	fmt.Printf("\r")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	core.Flushdb()
	core.StopTicker(ticker, done)

	time.Sleep(1 * time.Second) // Ensure everything is clean before exiting
	fmt.Printf("exit\n")
	os.Exit(0)
}
