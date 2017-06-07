package utils

type NestedEntityError struct {
    InnerError error
    Code 		   int
}

func (e NestedEntityError) Error() string {
	return e.InnerError.Error()
}
