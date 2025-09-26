package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	lock sync.RWMutex
	data map[string][]byte
}

func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Get(Key []byte) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	keyStr := string(Key)
	val, ok := c.data[keyStr]
	if !ok {
		return nil, fmt.Errorf("key (%s) not found", keyStr)
	}
	return val, nil
}

func (c *Cache) Set(Key []byte, Value []byte, ttl time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	ticker := time.NewTicker(ttl)
	go func() {
		<-ticker.C
		delete(c.data, string(Key))
	}()
	keyStr := string(Key)
	c.data[keyStr] = Value
	return nil
}

func (c *Cache) Has(Key []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, ok := c.data[string(Key)]
	return ok
}

func (c *Cache) Delete(Key []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, string(Key))
	return nil
}
