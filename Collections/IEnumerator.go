package Collections

// Supports a simple iteration over a collection.
type IEnumerator[T any] interface {

	// Sets the enumerator to its initial position, which is before the first element in the collection.
	Reset()

	// Advances the enumerator to the next element of the collection.
	//
	// Returns:
	//	object T - The element in the collection at the current position of the enumerator.
	//	ok bool - true if the enumerator was successfully advanced to the next element; false if the enumerator has passed the end of the collection.
	GetNext() (object T, ok bool)

	// Gets the element in the collection at the current position of the enumerator.
	//
	// Returns:
	//	object T - The element in the collection at the current position of the enumerator.
	//GetCurrent() (object T)
}
