package api

type OutOfScopeError struct {
	// TODO: add fields to describe the error
}

func (e *OutOfScopeError) Error() string {
	// TODO: implement a meaningful error message
	return "the requested operation is out of scope"
}
