package Collections

type Stack[T any] []T

func (stack *Stack[T]) Clear() {
	*stack = Stack[T]{}
}

func (stack Stack[T]) Count() (count int) {
	return len(stack)
}

func (stack Stack[T]) Peek() (item T, err error) {
	if len(stack) > 0 {
		item = (stack)[0]
		return item, err
	}
	return item, ErrCollectionContainsNoElements
}

func (stack *Stack[T]) Push(item T) {
	*stack = append(*stack, item)
}

func (stack *Stack[T]) Pop() (item T, err error) {
	if len(*stack) > 0 {
		item = (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		return item, err
	}
	return item, ErrCollectionContainsNoElements
}

func (stack Stack[T]) GetEnumerator() IEnumerator[T] {
	return &enumerator[T]{
		index:  0,
		object: stack,
	}
}
