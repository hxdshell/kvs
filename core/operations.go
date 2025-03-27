package core

import (
	"log"
	"strconv"
	"sync"
)

var store map[string][]byte
var rwlock sync.RWMutex

func InitStore() {
	store = make(map[string][]byte)
}

func Get(key string) []byte {
	rwlock.RLock()
	val := store[key]
	rwlock.RUnlock()
	return val
}

func Set(key string, value []byte) {
	rwlock.Lock()
	store[key] = value
	rwlock.Unlock()
}

func IncDec(key string, magnitude int, isInc bool) error {
	rwlock.Lock()
	value := store[key]
	val, err := strconv.Atoi(string(value))

	if err != nil {
		rwlock.Unlock()
		return err
	}

	if isInc {
		store[key] = []byte(strconv.Itoa(val + magnitude))
	} else {
		store[key] = []byte(strconv.Itoa(val - magnitude))
	}
	rwlock.Unlock()

	return nil
}

func List() [][]string {
	rwlock.RLock()

	var result [][]string
	for k, v := range store {
		result = append(result, []string{k, string(v)})
	}

	rwlock.RUnlock()
	log.Printf("LIST(%d)\n", len(result))
	return result
}

func Remove(key string) {
	rwlock.Lock()
	delete(store, key)
	rwlock.Unlock()
}

func Flushdb() {
	rwlock.Lock()
	clear(store)
	rwlock.Unlock()
}
