package gedis

import "time"

type DBGetFunc func() string

type SimpleCache struct {
	Operation *StringOperation
	Expire    time.Duration
	Getter    DBGetFunc
}

func NewSimpleCache(operation *StringOperation, expire time.Duration) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire}
}

func (c *SimpleCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

func (c *SimpleCache) GetCache(key string) (v interface{}) {
	v = c.Operation.Get(key).UnwrapOrElse(c.Getter)
	c.SetCache(key, v)
	return
}
