package Collections

type IEnumerable[T any] interface {

	// Returns an enumerator that iterates through a collection.
	//
	// Returns:
	//	enumerator IEnumertor[T] - An IEnumerator object that can be used to iterate through the collection.
	GetEnumerator() (enumerator IEnumerator[T])
}
