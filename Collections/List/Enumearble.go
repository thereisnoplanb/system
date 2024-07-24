package List

import "github.com/golangframework/system/Collections"

func (list *List[T]) GetEnumerator() Collections.IEnumerator[T] {
	return &enumerator[T]{
		index:  -1,
		object: list,
	}
}
