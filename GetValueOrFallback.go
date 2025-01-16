package system

// Returns the value pointed to by the pointer if the pointer is not nil; otherwise fallback value.
//
// Parameters:
//
//	pointer *T - Pointer to some value.
//	fallback T - Value returned when pointer points to nil
//
// Returns:
//
//	result - The value pointed to by the pointer or fallback value if the pointer points to nil.
func GetValueOrFallback[T any](pointer *T, fallback T) (result T) {
	if pointer != nil {
		return *pointer
	}
	return fallback
}
