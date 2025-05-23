package exercise10

type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](s []T) *List[T] {
	if len(s) == 0 {
		return nil
	}

	head := &List[T]{nil, s[0]}
	prev := head

	for _, v := range s[1:] {
		next := &List[T]{nil, v}
		prev.next = next
		prev = next
	}

	return head
}
