package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var Store map[string][]byte
var rwlock sync.RWMutex

func get(key string) []byte {
	rwlock.RLock()
	val := Store[key]
	rwlock.RUnlock()
	return val
}

func put(key string, value []byte) {
	rwlock.Lock()
	Store[key] = value
	rwlock.Unlock()
}

func main() {

	Store = make(map[string][]byte)
	port := 3000
	addr := fmt.Sprintf(":%d", port)

	server := &http.Server{
		Addr:    addr,
		Handler: GetMux(),
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Could not start the server at port %d\n", port)
			os.Exit(1)
		}
	}()
	fmt.Printf("Starting a server on port %d\n", port)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	clear(Store)
	fmt.Println()
	server.Shutdown(context.Background())
}
