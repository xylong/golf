package gedis

import "time"

const Expire = "expire"

type OperationAttr struct {
	Name  string
	Value interface{}
}

type OperationAttrs []*OperationAttr

func (o OperationAttrs) Find(name string) interface{} {
	for _, attr := range o {
		if attr.Name == name {
			return attr.Value
		}
	}

	return nil
}

// WithExpire 设置过期时间
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{
		Name:  Expire,
		Value: t,
	}
}
