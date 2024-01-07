package hashmap

import linkedlist "gostl/linked_list"

type hashable interface {
	int | int32 | int64 | string
}

type Hashmap[T hashable, Z any] struct {
	TableSize int
	Values    []*linkedlist.LinkedList[any]
}

func NewHashMap[T hashable, Z any](size int) *Hashmap[T, Z] {
	v := make([]*linkedlist.LinkedList[any], size)
	h := &Hashmap[T, Z]{
		TableSize: size,
		Values:    v,
	}
	return h
}

func (h *Hashmap[T, Z]) Put(key T, value Z) {
	index := Hash(string(key))
	h.Values[index].Add(value)
}

func (h *Hashmap[T, Z]) Get() {}

func (h *Hashmap[T, Z]) Remove() {}

func (h *Hashmap[T, Z]) Contains() {}

func (h *Hashmap[T, Z]) Size() {}
