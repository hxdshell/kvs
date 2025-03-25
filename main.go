package main

import (
	"context"
	"fmt"
	"log"
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

func list() [][]string {
	rwlock.RLock()

	var result [][]string
	for k, v := range Store {
		result = append(result, []string{k, string(v)})
	}

	rwlock.RUnlock()
	log.Printf("LIST(%d)\n", len(result))
	return result
}

func set(key string, value []byte) {
	rwlock.Lock()
	Store[key] = value
	rwlock.Unlock()
}

func remove(key string) {
	rwlock.Lock()
	delete(Store, key)
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
	fmt.Printf("\033[1;37mStarting a server on port %d\033[0m\n\n", port)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	clear(Store)
	fmt.Printf("\rexit\n")
	server.Shutdown(context.Background())
}
