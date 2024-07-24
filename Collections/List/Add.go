package List

func (list *List[T]) Add(item T) {
	(*list) = append((*list), item)
}
