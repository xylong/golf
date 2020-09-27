package gedis

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"
)

const (
	SerializeJson = "json"
	SerializeGOB  = "gob"
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
	obj := c.Getter()

	switch c.Serializer {
	case SerializeJson:
		f := func() string {
			if bytes, err := json.Marshal(obj); err != nil {
				return ""
			} else {
				return string(bytes)
			}
		}
		v = c.Operation.Get(key).UnwrapOrElse(f)
	case SerializeGOB:
		f := func() string {
			buf := &bytes.Buffer{}
			encode := gob.NewEncoder(buf)
			if err := encode.Encode(obj); err != nil {
				return ""
			} else {
				return buf.String()
			}
		}
		v = c.Operation.Get(key).UnwrapOrElse(f)
	}

	c.SetCache(key, v)
	return
}

func (c *SimpleCache) GetCacheForObject(key string, obj interface{}) interface{} {
	result := c.GetCache(key)
	if result == nil {
		return nil
	}

	switch c.Serializer {
	case SerializeJson:
		if err := json.Unmarshal([]byte(result.(string)), obj); err != nil {
			return nil
		}
	case SerializeGOB:
		buff := &bytes.Buffer{}
		buff.WriteString(result.(string))
		decode := gob.NewDecoder(buff)
		if decode.Decode(obj) != nil {
			return nil
		}
	}
	return nil
}
