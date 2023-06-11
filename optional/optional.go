package optional

import (
	"fmt"

	"github.com/tomasbasham/gonads"
)

// Func is a function that takes a value and returns an optional value.
type Func[T, U any] Optional[gonads.Func[T, U]]

// Optional is a container for a value that may or may not be present. It may
// be in either one of two states:
//
//   - Some: the value is present in the container.
//   - None: the value is absent from the container.
type Optional[T any] interface {
	Unwrap() (T, error)
}

// Some is a container for a value that is present.
type Some[T any] struct {
	value T
}

// Pure returns a Some[T] containing the given value.
func Pure[T any](value T) Some[T] {
	return Some[T]{value}
}

// NewSome is an alias for Pure.
func NewSome[T any](value T) Some[T] {
	return Pure(value)
}

// Unwrap returns the value contained in a Some[T].
func (o Some[T]) Unwrap() (T, error) {
	return o.value, nil
}

// None is a container for a value that is absent.
type None[T any] struct{}

// NewNone returns a None[T].
func NewNone[T any]() None[T] {
	return None[T]{}
}

// Unwrap returns the value contained in a None[T].
func (o None[T]) Unwrap() (T, error) {
	return *new(T), fmt.Errorf("failed to unwrap None")
}

// Map maps a function over an optional value.
func Map[T, U any](o Optional[T], f gonads.Func[T, U]) Optional[U] {
	if s, ok := o.(Some[T]); ok {
		return Pure(f(s.value))
	}
	return None[U]{}
}

// Apply applies an optional function to an optional value.
func Apply[T, U any](o Optional[T], f Func[T, U]) Optional[U] {
	if f, ok := f.(Some[gonads.Func[T, U]]); ok {
		return Map(o, f.value)
	}
	return None[U]{}
}

// FlatMap flat-maps a function over an optional value.
func FlatMap[T, U any](o Optional[T], f gonads.Func[T, Optional[U]]) Optional[U] {
	if s, ok := o.(Some[T]); ok {
		return f(s.value)
	}
	return None[U]{}
}
