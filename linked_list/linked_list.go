package linkedlist

import "fmt"

type node[T comparable] struct {
	Value T
	Next  *node[T]
}

type LinkedList[T comparable] struct {
	Head *node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	ll := &LinkedList[T]{
		Head: nil,
	}
	return ll
}

func newNode[T comparable](val T) *node[T] {
	nn := &node[T]{
		Value: val,
		Next:  nil,
	}

	return nn
}

func (ll *LinkedList[T]) Add(element T) {
	nn := newNode[T](element)
	fmt.Println("ll")
	fmt.Println(ll)
	if ll.Head == nil {
		// This might bu useless bro
		ll.Head = nn
		return
	}

	curNodePtr := *ll.Head

	for {
		if curNodePtr.Next == nil {
			curNodePtr.Next = nn
			return
		}

		curNodePtr = *curNodePtr.Next
	}
}

func (ll *LinkedList[T]) Get(element T) (T, bool) {
	curNodePtr := ll.Head

	for {
		fmt.Println(ll)
		fmt.Println(element)
		el := any(element).(T)
		/*
			if !ok {
				fmt.Println("case element failed")
			}
			val, ok := any(curNodePtr.Value).(T)
			if !ok {
				fmt.Println("cast value failed")
			}
		*/
		if curNodePtr.Value == el {
			return el, true
		}

		if curNodePtr.Next == nil {
			return el, false
		}

		curNodePtr = curNodePtr.Next
	}
}
