package reader

import "github.com/tomasbasham/gonads"

// Reader is a function that takes an environment and returns a value. It is
// used to model computations that depend on some global state. It is also
// known as the environment monad.
type Reader[E, A any] struct {
	g func(E) A
}

// Run runs a reader and returns its value.
func (r *Reader[E, A]) Run(env E) A {
	return r.g(env)
}

// Pure returns a Reader[E, A] containing the given value.
func Pure[E, A any](a A) Reader[E, A] {
	return Reader[E, A]{func(_ E) A { return a }}
}

// Map maps a function over a reader.
func Map[E, A, B any](r Reader[E, A], f gonads.Func[A, B]) Reader[E, B] {
	return Reader[E, B]{func(e E) B { return f(r.g(e)) }}
}

// Apply applies a reader function to a reader value.
func Apply[E, A, B any](r Reader[E, A], f Reader[E, gonads.Func[A, B]]) Reader[E, B] {
	return Reader[E, B]{func(e E) B { return f.g(e)(r.g(e)) }}
}

// FlatMap maps a function over a reader.
func FlatMap[E, A, B any](r Reader[E, A], f gonads.Func[A, Reader[E, B]]) Reader[E, B] {
	return Reader[E, B]{func(e E) B { return f(r.g(e)).g(e) }}
}

// Zip combines two readers into one.
func Zip[E, A, B, U any](ra Reader[E, A], rb Reader[E, B], f func(A, B) U) Reader[E, U] {
	return FlatMap(ra, func(a A) Reader[E, U] {
		return Map(rb, func(b B) U {
			return f(a, b)
		})
	})
}

// RunReader runs a reader and returns its value.
func RunReader[E, A any](r Reader[E, A], e E) A {
	return r.Run(e)
}
