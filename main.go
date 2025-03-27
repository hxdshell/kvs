package main

import (
	"context"
	"fmt"
	"kvs/core"
	"kvs/handlers"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	port := 3000
	addr := fmt.Sprintf(":%d", port)

	server := &http.Server{
		Addr:    addr,
		Handler: handlers.GetMux(),
	}

	core.InitStore()

	fmt.Printf("\033[1;37mStarting a server on port %d\033[0m\n\n", port)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Could not start the server at port %d\n", port)
			return
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	core.Flushdb()
	fmt.Printf("\rexit\n")
	server.Shutdown(context.Background())

	os.Exit(0)
}
