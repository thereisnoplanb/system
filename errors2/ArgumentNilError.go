package errors

type ArgumentNilError struct {
	name    string
	message string
}

func NewArgumentNilError(name string, message ...string) error {
	return ArgumentNilError{
		name:    name,
		message: getString(message),
	}
}

func (err ArgumentNilError) Error() string {

	result := "value cannot be nil"
	if err.name != "" {
		result += ", name: " + err.name
	}
	if err.message != "" {
		result += ", message: " + err.message
	}
	return result
}
