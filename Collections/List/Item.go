package List

func (list *List[T]) Item(index int) (item T) {
	return (*list)[index]
}
