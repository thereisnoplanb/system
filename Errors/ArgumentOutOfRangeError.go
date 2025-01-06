package Errors

type ArgumentOutOfRangeError struct {
	argumentName string
	message      string
}

func NewArgumentOutOfRangeError(argumentName string, message ...string) error {
	return ArgumentOutOfRangeError{
		argumentName: argumentName,
		message:      getString(message),
	}
}

func (err ArgumentOutOfRangeError) Error() string {

	result := "argument out of range"
	if err.argumentName != "" {
		result += ", argument name: " + err.argumentName
	}
	if err.message != "" {
		result += ", message: " + err.message
	}
	return result
}
