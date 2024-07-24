package Collections

type List[T any] []T

func (list *List[T]) Add(item T) {
	(*list) = append((*list), item)
}

func (list *List[T]) AddRange(collection IEnumerable[T]) {
	enumerator := collection.GetEnumerator()
	item, ok := enumerator.GetNext()
	for ok {
		list.Add(item)
		item, ok = enumerator.GetNext()
	}
}

func (list *List[T]) Capacity() (capacity int) {
	return cap(*list)
}

func (list *List[T]) Clear() {
	*list = List[T]{}
}

func (list *List[T]) Insert(index int, item T) (err error) {
	if index < 0 {
		return ErrIndexIsLessThanZero
	}
	if index > len(*list) {
		return ErrCollectionContainsNoElements
	}
	temp := append((*list)[:index], item)
	(*list) = append(temp, (*list)[index:]...)
	return nil
}

func (list *List[T]) RemoveAt(index int) (err error) {
	if index < 0 {
		return ErrIndexIsLessThanZero
	}
	if index >= len(*list) {
		return ErrIndexIsGreaterThanCount
	}
	(*list) = append((*list)[:index], (*list)[index+1:]...)
	return nil
}

func (list *List[T]) Item(index int) (item T) {
	return (*list)[index]
}

func (list *List[T]) InsertRange(index int, collection IEnumerable[T]) (err error) {
	if index < 0 {
		return ErrIndexIsLessThanZero
	}
	if index > len(*list) {
		return ErrCollectionContainsNoElements
	}
	enumerator := collection.GetEnumerator()
	item, ok := enumerator.GetNext()
	if !ok {
		return nil
	}
	temp := append((*list)[:index], item)
	for ok {
		temp.Add(item)
		item, ok = enumerator.GetNext()
	}
	(*list) = append(temp, (*list)[index:]...)
	return nil
}

func (list List[T]) GetEnumerator() IEnumerator[T] {
	return &enumerator[T]{
		index:  -1,
		object: list,
	}
}
