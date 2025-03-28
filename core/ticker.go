package core

import (
	"fmt"
	"time"
)

type Job func()

// Takes interval (in seconds) and performs the operation repeatedly at given intervals
func StartTicker(interval int, job Job) (*time.Ticker, chan bool) {
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopping recurring operation")
				return
			case t := <-ticker.C:
				fmt.Println("Performing operation at ", t)
				job()
			}
		}
	}()
	return ticker, done
}
func StopTicker(ticker *time.Ticker, done chan bool) {
	done <- true
	ticker.Stop()
}
