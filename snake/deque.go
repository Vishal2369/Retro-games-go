package main

type Deque[T any] struct {
	head *node[T]
	tail *node[T]
	len  int32
}

type node[T any] struct {
	val      T
	nextNode *node[T]
	prevNode *node[T]
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (d *Deque[T]) Len() int32 {
	return d.len
}

func (d *Deque[T]) IsEmpty() bool {
	return d.len == 0
}

func (d *Deque[T]) PushFront(elem T) {
	newNode := &node[T]{val: elem, nextNode: nil, prevNode: nil}

	if d.len == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		d.head.nextNode = newNode
		newNode.prevNode = d.head
		d.head = newNode
	}

	d.len++
}

func (d *Deque[T]) PushBack(elem T) {
	newNode := &node[T]{val: elem, nextNode: nil, prevNode: nil}

	if d.len == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.nextNode = d.tail
		d.tail.prevNode = newNode
		d.tail = newNode
	}

	d.len++
}

func (d *Deque[T]) PopFront() T {
	if d.IsEmpty() {
		panic("deque is already empty")
	}

	elem := d.head.val

	if d.len == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.prevNode
        d.head.nextNode = nil
	}

	d.len--

	return elem
}

func (d *Deque[T]) PopBack() T {
	if d.IsEmpty() {
		panic("deque is already empty")
	}

	elem := d.tail.val

	if d.len == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.nextNode
        d.tail.prevNode = nil
	}

	d.len--

	return elem
}

func (d *Deque[T]) Front() T {
    if d.IsEmpty() {
        panic("deque is already empty")
    }

    return d.head.val;
}

func (d *Deque[T]) AllElements() []T {
	var elems []T
	for s := d.head; s != nil; s = s.prevNode {
		elems = append(elems, s.val)
	}
	return elems
}
