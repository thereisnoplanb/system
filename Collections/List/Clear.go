package List

func (list *List[T]) Clear() {
	*list = List[T]{}
}
