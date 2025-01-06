package errors

type FormatError struct {
	argumentName string
	message      string
}

func NewFormatError(argumentName string, message ...string) error {
	return FormatError{
		argumentName: argumentName,
		message:      getString(message),
	}
}

func (err FormatError) Error() string {

	result := "format"
	if err.argumentName != "" {
		result += ", argument name: " + err.argumentName
	}
	if err.message != "" {
		result += ", message: " + err.message
	}
	return result
}
