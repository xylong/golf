package gedis

import (
	"context"
	"time"
)

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

// Set
func (s *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *StringResult {
	expiration := OperationAttrs(attrs).Find(Expire)
	if expiration == nil {
		expiration = 0
	}
	return NewStringResult(Redis().Set(s.ctx, key, value, expiration.(time.Duration)).Result())
}

// Get
func (s *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(s.ctx, key).Result())
}

func (s *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(s.ctx, keys...).Result())
}
