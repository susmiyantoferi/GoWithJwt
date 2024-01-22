package exception

type ErrorUnauthorized struct {
	Error string
}

func NewErrorUnauthorized(error string) *ErrorUnauthorized {
	return &ErrorUnauthorized{Error: error}
}
