package hashmap

import (
	"errors"
	linkedlist "gostl/linked_list"
)

type Hashable interface {
	int | int32 | int64 | string
}

type Hashmap[T Hashable, Z any] struct {
	TableSize int
	Values    []*linkedlist.LinkedList[any]
}

func NewHashMap[T Hashable, Z any](size int) *Hashmap[T, Z] {
	v := make([]*linkedlist.LinkedList[any], size)
	h := &Hashmap[T, Z]{
		TableSize: size,
		Values:    v,
	}
	return h
}

func (h *Hashmap[T, Z]) Put(key T, value Z) {
	index := Hash(key)
	h.Values[index].Add(value)
}

func (h *Hashmap[T, Z]) Get(key T) (*Z, error) {
	index := Hash(key)
	el, ok := h.Values[index].Get(key)
	if !ok {
		return nil, errors.New("couldn't get the element")
	}
	casted, ok := el.(Z)
	if !ok {
		return nil, errors.New("couldn't case type")
	}

	return &casted, nil
}

func (h *Hashmap[T, Z]) Remove() {}

func (h *Hashmap[T, Z]) Contains() {}

func (h *Hashmap[T, Z]) Size() {}
