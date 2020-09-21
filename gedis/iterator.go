package gedis

// Iterator 迭代器
type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

// HasNext 判断是否有后续值
func (i *Iterator) HasNext() bool {
	if length := len(i.data); i.data == nil || length == 0 {
		return false
	} else {
		return i.index < length
	}
}

// Next 下一个
func (i *Iterator) Next() interface{} {
	value := i.data[i.index]
	i.index = i.index + 1
	return value
}
