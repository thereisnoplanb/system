package List

func (list *List[T]) Capacity() (capacity int) {
	return cap(*list)
}
