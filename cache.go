package cache

import "time"

type Cache struct {
	m map[string]string
}

func NewCache() Cache {
	return Cache{make(map[string]string)}
}

func (c Cache) Get(key string) (string, bool) {
	s, ok := c.m[key]
	return s, ok
}

func (c *Cache) Put(key, value string) {
	c.m[key] = value
}

func (c Cache) Keys() []string {
	var result []string
	for _, vl := range c.m {
		result = append(result, vl)
	}
	return result
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
}
