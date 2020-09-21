package gedis

import "context"

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (s *StringOperation) Set() {

}

func (s *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(s.ctx, key).Result())
}

func (s *StringOperation) MGet(keys ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(s.ctx, keys...).Result())
}
