package gedis

import (
	"encoding/json"
	"time"
)

const (
	SerializeJson = "json"
)

type DBGetFunc func() interface{}

type SimpleCache struct {
	Operation  *StringOperation // 操作类
	Expire     time.Duration    // 过期时间
	Getter     DBGetFunc        // db
	Serializer string           // 序列化方式
}

func NewSimpleCache(operation *StringOperation, expire time.Duration, serializer string) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire, Serializer: serializer}
}

func (c *SimpleCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

func (c *SimpleCache) GetCache(key string) (v interface{}) {
	switch c.Serializer {
	case SerializeJson:
		f := func() string {
			obj := c.Getter()
			if bytes, err := json.Marshal(obj); err != nil {
				return ""
			} else {
				return string(bytes)
			}
		}
		v = c.Operation.Get(key).UnwrapOrElse(f)
	}

	c.SetCache(key, v)
	return
}
