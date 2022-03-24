package errors

type basicError struct {
	cause string
	err   string
}

func (e *basicError) Error() string {

	result := e.err
	if e.cause != "" {
		result = result + ": " + e.cause
	}
	return result
}

func (e *basicError) SetCause(cause string) *basicError {

	e.cause = cause
	return e
}

func new(err string) *basicError {
	return &basicError{
		err: err,
	}
}
