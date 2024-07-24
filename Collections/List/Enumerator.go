package List

type enumerator[T any] struct {
	index  int
	object *List[T]
}

func (enumerator *enumerator[T]) Reset() {
	enumerator.index = -1
}

func (enumerator *enumerator[T]) GetNext() (object T, ok bool) {
	if enumerator.index < len(*enumerator.object) {
		object = (*enumerator.object)[enumerator.index]
		enumerator.index++
		ok = true
	}
	return object, ok
}
