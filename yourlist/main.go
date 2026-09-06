package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Append(item T) {
	for l.next != nil {
		l = l.next
	}
	l.next = &List[T]{val: item}
}

func main() {

	l := &List[int]{}

	l.Append(2)

	println(l.val)

}
