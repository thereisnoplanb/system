package system

func GetPointer[T any](value T) (pointer *T) {
	return &value
}
