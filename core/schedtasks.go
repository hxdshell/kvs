package core

import (
	"fmt"
	"time"
)

func KillExpiredKeys() {
	for key, data := range store {
		if data.ttl.IsZero() {
			continue
		}

		if time.Now().After(data.ttl) {
			fmt.Println("Deleting ", key)
			delete(store, key)
		}
	}
}
