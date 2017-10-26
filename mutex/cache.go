package main

import (
	"fmt"
	"sync"
	"time"
)

type entry struct {
	value     string
	expiresAt time.Time
}

//Cache is a TTL cache that is safe for concurrent use
type Cache struct {
	entries map[string]*entry
	//TODO: protect this for concurrent use!
	mx sync.RWMutex
}

//NewCache constructs a new Cache object
func NewCache() *Cache {
	c := &Cache{
		entries: map[string]*entry{},
	}

	//start janitor on separate goroutine
	go c.janitor()

	return c
}

//Set adds a key/value to the cache
func (c *Cache) Set(key string, value string, timeToLive time.Duration) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.entries[key] = &entry{value, time.Now().Add(timeToLive)}
}

//Get gets the value associated with a key
func (c *Cache) Get(key string) (string, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	entry, found := c.entries[key]
	if !found {
		return "", false
	}
	return entry.value, true
}

//janitor removes expired nodes
func (c *Cache) janitor() {
	for {
		//only needs to clean every so often
		time.Sleep(time.Second)

		now := time.Now()

		//clean removes nodes (mutates state)
		//must lock and unlock node
		c.mx.Lock()

		fmt.Printf("janitor is running")
		for key, entry := range c.entries {
			if entry.expiresAt.Before(now) {
				fmt.Printf("purging key %s/n", key)
				delete(c.entries, key)
			}
		}

		//can't use defer because this function never exits
		//manually call unlock
		c.mx.Unlock()
	}
}
