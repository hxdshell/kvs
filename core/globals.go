package core

import (
	"sync"
	"time"
)

// Some structures
type Value struct {
	val []byte
	ttl time.Time
}

// One file to keep track of all singleton structures
var store map[string]Value
var rwlock sync.RWMutex
