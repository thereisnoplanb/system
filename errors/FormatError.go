package errors

type FormatError struct {
	name    string
	message string
}

func NewFormatError(name string, message ...string) error {
	return FormatError{
		name:    name,
		message: getString(message),
	}
}

func (err FormatError) Error() string {

	result := "invalid format"
	if err.name != "" {
		result += ", name: " + err.name
	}
	if err.message != "" {
		result += ", message: " + err.message
	}
	return result
}
