package errors

type ArgumentOutOfRangeError struct {
	name    string
	message string
}

func NewArgumentOutOfRangeError(name string, message ...string) error {
	return ArgumentOutOfRangeError{
		name:    name,
		message: getString(message),
	}
}

func (err ArgumentOutOfRangeError) Error() string {

	result := "value out of range"
	if err.name != "" {
		result += ", name: " + err.name
	}
	if err.message != "" {
		result += ", message: " + err.message
	}
	return result
}
