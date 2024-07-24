package List

func (list *List[T]) Count() (count int) {
	return len(*list)
}
