package cache

import "time"

type Cache struct {
	m                               map[string]expiringCache
	amountExpiringCacheWithDeadline int
	isRunning                       bool
}

type expiringCache struct {
	value    string
	deadline time.Time
}

func (c *Cache) deleteExpired() {
	for key, value := range c.m {
		if time.Now().Before(value.deadline) {
			delete(c.m, key)
		}
	}
}

func (c *Cache) waitExpiringCache() {
	for {
		time.Sleep(1 * time.Second)
		c.deleteExpired()
		c.amountExpiringCacheWithDeadline--
		if c.amountExpiringCacheWithDeadline == 0 {
			c.isRunning = false
			break
		}
	}
}

func NewCache() Cache {
	cache := Cache{m: make(map[string]expiringCache)}
	return cache
}

func (c Cache) Get(key string) (string, bool) {
	s, ok := c.m[key]
	return s.value, ok
}

func (c *Cache) Put(key, value string) {
	c.m[key] = expiringCache{value: value}
}

func (c Cache) Keys() []string {
	var result []string
	for key := range c.m {
		result = append(result, key)
	}
	return result
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.m[key] = expiringCache{value: value, deadline: deadline}
	if !c.isRunning {
		c.amountExpiringCacheWithDeadline++
		c.waitExpiringCache()
		c.isRunning = true
	}
}
