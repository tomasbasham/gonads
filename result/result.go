package result

import "github.com/tomasbasham/gonads"

// Func is a function that takes a value and returns a result value.
type Func[T, U any] Result[gonads.Func[T, U]]

// Result is a container for a value that is the result of some operation. It
// may be in either one of two states:
//
//   - Success: the operation was successful and the result value is present in
//     the container.
//   - Failure: the operation failed and the error value is present in the
//     container.
type Result[T any] interface {
	Unwrap() (T, error)
}

// Success is a container for a value that is the result of a successful
// operation.
type Success[T any] struct {
	value T
}

// Pure returns a Success[T] containing the given value.
func Pure[T any](value T) Success[T] {
	return Success[T]{value}
}

// NewSuccess is an alias for Pure.
func NewSuccess[T any](value T) Success[T] {
	return Pure(value)
}

// Unwrap returns the value contained in a Success[T].
func (r Success[T]) Unwrap() (T, error) {
	return r.value, nil
}

// Failure is a container for a value that is the result of a failed operation.
type Failure[T any] struct {
	value error
}

// NewFailure returns a Failure[T] containing the given error.
func NewFailure[T any](value error) Failure[T] {
	return Failure[T]{value}
}

// Unwrap returns the value contained in a Failure[T].
func (r Failure[T]) Unwrap() (T, error) {
	return *new(T), r.value
}

// Map maps a function over a result value.
func Map[T, U any](r Result[T], f gonads.Func[T, U]) Result[U] {
	if s, ok := r.(Success[T]); ok {
		return Pure(f(s.value))
	}

	_, err := r.Unwrap()
	return Failure[U]{err}
}

// Apply applies a result function to a result value.
func Apply[T, U any](r Result[T], f Func[T, U]) Result[U] {
	if f, ok := f.(Success[gonads.Func[T, U]]); ok {
		return Map(r, f.value)
	}

	_, err := r.Unwrap()
	return Failure[U]{err}
}

// FlatMap flat-maps a function over a result value.
func FlatMap[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if s, ok := r.(Success[T]); ok {
		return f(s.value)
	}

	_, err := r.Unwrap()
	return Failure[U]{err}
}
