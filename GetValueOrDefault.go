package system

func GetValueOrDefault[T any](pointer *T) (result T) {
	if pointer != nil {
		result = *pointer
	}
	return result
}
