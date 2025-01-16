package system

// Returns the value pointed to by the pointer if the pointer is not nil; otherwise default value.
//
// Parameters:
//
//	pointer *T - Pointer to some value.
//
// Returns:
//
//	result - The value pointed to by the pointer or default value if the pointer points to nil.
func GetValueOrDefault[T any](pointer *T) (result T) {
	if pointer != nil {
		result = *pointer
	}
	return result
}
