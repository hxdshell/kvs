package core

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func InitStore() {
	store = make(map[string]Value)
}

func Get(key string) []byte {
	var result []byte = []byte{}

	rwlock.RLock()
	data := store[key]
	if time.Now().Before(data.ttl) {
		result = data.val
	}
	rwlock.RUnlock()

	return result
}

func Set(key string, storeval []byte, ttl int) {

	fmt.Println("TTL INT : ", ttl)
	var storeTTL time.Time
	if ttl < 0 {
		storeTTL = time.Time{}
	} else {
		fmt.Println("TTL DURATION: ", time.Duration(ttl))
		storeTTL = time.Now().Add(time.Duration(ttl) * time.Second)
		fmt.Println("storeTTL ", storeTTL)
	}

	rwlock.Lock()
	store[key] = Value{
		val: storeval,
		ttl: storeTTL,
	}
	rwlock.Unlock()
	fmt.Printf("%02d-%02d-%02d\n", store[key].ttl.Hour(), store[key].ttl.Minute(), store[key].ttl.Second())
}

func IncDec(key string, magnitude int, isInc bool) error {
	rwlock.Lock()
	value := store[key].val
	val, err := strconv.Atoi(string(value))

	if err != nil {
		rwlock.Unlock()
		return err
	}

	prevVal := store[key]
	if isInc {
		prevVal.val = []byte(strconv.Itoa(val + magnitude))
	} else {
		prevVal.val = []byte(strconv.Itoa(val - magnitude))
	}
	store[key] = prevVal
	rwlock.Unlock()

	return nil
}

func List() [][]string {
	rwlock.RLock()

	var result [][]string
	for k, v := range store {
		if time.Now().Before(v.ttl) {
			result = append(result, []string{k, string(v.val)})
		}
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
