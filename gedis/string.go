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
func (s *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *InterfaceResult {
	expiration := OperationAttrs(attrs).Find(Expire)

	if nx := OperationAttrs(attrs).Find(NX).UnwrapOr(nil); nx != nil {
		return NewInterfaceResult(Redis().SetNX(s.ctx, key, value, expiration.UnwrapOr(time.Second*0).(time.Duration)).Result())
	}

	if xx := OperationAttrs(attrs).Find(XX).UnwrapOr(nil); xx != nil {
		return NewInterfaceResult(Redis().SetXX(s.ctx, key, value, expiration.UnwrapOr(time.Second*0).(time.Duration)).Result())
	}

	return NewInterfaceResult(Redis().Set(s.ctx, key, value, expiration.UnwrapOr(time.Second*0).(time.Duration)).Result())
}

// Get
func (s *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(s.ctx, key).Result())
}

func (s *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(s.ctx, keys...).Result())
}
