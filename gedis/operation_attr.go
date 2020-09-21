package gedis

import (
	"fmt"
	"time"
)

const (
	Expire = "expire"
	Nx     = "nx"
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

func WithNx() *OperationAttr {
	return &OperationAttr{
		Name:  Nx,
		Value: struct{}{},
	}
}
