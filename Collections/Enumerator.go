package Collections

type keyValueEnumerator[TKey comparable, TValue any] struct {
	Index  int
	Object map[TKey]TValue
	Keys   []TKey
}

func (enumerator *keyValueEnumerator[TKey, TValue]) Reset() {
	enumerator.Index = 0
}

func (enumerator *keyValueEnumerator[TKey, TValue]) GetNext() (object KeyValue[TKey, TValue], ok bool) {
	if enumerator.Index < len(enumerator.Object) {
		key := enumerator.Keys[enumerator.Index]
		object = KeyValue[TKey, TValue]{
			Key:   key,
			Value: (enumerator.Object)[key],
		}
		enumerator.Index++
		ok = true
	}
	return object, ok
}

type enumerator[T any] struct {
	index  int
	object []T
}

func (enumerator *enumerator[T]) Reset() {
	enumerator.index = -1
}

func (enumerator *enumerator[T]) GetNext() (object T, ok bool) {
	if enumerator.index < len(enumerator.object) {
		object = enumerator.object[enumerator.index]
		enumerator.index++
		ok = true
	}
	return object, ok
}
