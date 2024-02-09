package lib

const MAX_STACK_SIZE = 255

type Stack[T interface{}] struct {
	items        []T
	stackPointer int
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		items:        []T{},
		stackPointer: -1,
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
	s.stackPointer++
	if len(s.items) > MAX_STACK_SIZE {
		panic("stack overflow")
	}
}

func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.stackPointer--
	return item
}

func (s *Stack[T]) Top() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Clear() {
	s.items = []T{}
}
