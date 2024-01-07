package linkedlist

type node[T comparable] struct {
	Value T
	Next  *node[T]
}

type LinkedList[T comparable] struct {
	Head *node[T]
}

func NewLinkedList[T comparable]() LinkedList[T] {
	ll := LinkedList[T]{}
	return ll
}

func newNode[T comparable](val T) node[T] {
	nn := node[T]{
		Value: val,
		Next:  nil,
	}

	return nn
}

func (ll *LinkedList[T]) Add(element T) {
	nn := newNode[T](element)
	if ll.Head == nil {
		ll.Head = &nn
		return
	}

	curNodePtr := *ll.Head

	for {
		if curNodePtr.Next == nil {
			curNodePtr.Next = &nn
			return
		}

		curNodePtr = *curNodePtr.Next
	}
}

func (ll *LinkedList[T]) Get(element T) (T, bool) {
	curNodePtr := *ll.Head

	for {
		el := any(element).(T)
		val := any(curNodePtr.Value).(T)

		if val == el {
			return el, true
		}

		if curNodePtr.Next == nil {
			return el, false
		}

		curNodePtr = *curNodePtr.Next
	}
}
