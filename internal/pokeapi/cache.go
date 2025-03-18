package pokeapi

import (
	"fmt"
	"sync"
	"time"
)

type pokeCache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func newPokeCache(interval time.Duration) *pokeCache {
	fmt.Println("// Creating new pokeCache")
	pc := &pokeCache{
		cacheMap: make(map[string]cacheEntry),
	}
	fmt.Println("// Starting reapLoop")
	go pc.reapLoop(interval) // run in a new goroutine
	return pc
}

func (pc *pokeCache) Add(key string, val []byte) {
	// adds a new entry to the cache
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.cacheMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (pc *pokeCache) Get(key string) ([]byte, bool) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	elem, ok := pc.cacheMap[key]
	if !ok {
		fmt.Println("err: entry not found")
		return nil, false
	}
	return elem.val, true
}

func (pc *pokeCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker((interval))
	defer ticker.Stop()

	for {
		<-ticker.C
		// lock when accessing the map, then unlock
		pc.mu.Lock()
		for key, elem := range pc.cacheMap {
			if time.Since(elem.createdAt) > interval {
				delete(pc.cacheMap, key)
			}
		}
		pc.mu.Unlock()
	}
}
