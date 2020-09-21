package gedis

import (
	"fmt"
	"time"
)

const (
	Expire = "expire"
	NX     = "nx"
	XX     = "xx"
)

var (
	empty = struct{}{}
)

// OperationAttr 属性操作
type OperationAttr struct {
	Name  string
	Value interface{}
}

type OperationAttrs []*OperationAttr

func (o OperationAttrs) Find(name string) *InterfaceResult {
	for _, attr := range o {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}

	return NewInterfaceResult(nil, fmt.Errorf("OperationAttrs found error: %s", name))
}

// WithExpire 设置过期时间
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{
		Name:  Expire,
		Value: t,
	}
}

func WithNX() *OperationAttr {
	return &OperationAttr{
		Name:  NX,
		Value: empty,
	}
}

func WithXX() *OperationAttr {
	return &OperationAttr{
		Name:  XX,
		Value: empty,
	}
}
