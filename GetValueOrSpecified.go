package system

func GetValueOrSpecified[T any](pointer *T, specified T) (result T) {
	if pointer != nil {
		return *pointer
	}
	return specified
}
