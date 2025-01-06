package system

func GetValueOrSpecified[T any](pointer *T, specified T) (result T) {
	if pointer != nil {
		result = *pointer
	} else {
		result = specified
	}
	return result
}
