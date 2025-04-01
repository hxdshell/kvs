package core

import (
	"time"
)

func KillExpiredKeys() {
	for key, data := range store {
		if data.ttl.IsZero() {
			continue
		}

		if time.Now().After(data.ttl) {
			delete(store, key)
		}
	}
}
