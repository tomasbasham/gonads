package writer

import "github.com/tomasbasham/gonads"

// Writer is a monad that writes a value of type A and an error of type E. It
// is used to model computations that may fail and produce a log.
type Writer[E, A any] struct {
	g func() (A, E)
}

// Run runs a writer and returns its value and error.
func (w *Writer[E, A]) Run() (A, E) {
	return w.g()
}

// Pure returns a Writer[E, A] containing the given value.
func Pure[E, A any](a A) Writer[E, A] {
	return Writer[E, A]{func() (A, E) { return a, *new(E) }}
}

// Map maps a function over a writer.
func Map[E, A, B any](w Writer[E, A], f gonads.Func[A, B]) Writer[E, B] {
	return Writer[E, B]{func() (B, E) {
		a, e := w.g()
		return f(a), e
	}}
}

// Apply applies a writer function to a writer value.
func Apply[E, A, B any](w Writer[E, A], f Writer[E, gonads.Func[A, B]]) Writer[E, B] {
	return Writer[E, B]{func() (B, E) {
		a, e := w.g()
		g, _ := f.g()
		return g(a), e
	}}
}

// FlatMap maps a function over a writer.
func FlatMap[E, A, B any](w Writer[E, A], f gonads.Func[A, Writer[E, B]]) Writer[E, B] {
	return Writer[E, B]{func() (B, E) {
		a, _ := w.g()
		return f(a).g()
	}}
}

// Zip combines two writers into one.
func Zip[E, A, B, U any](wa Writer[E, A], wb Writer[E, B], f func(A, B) U) Writer[E, U] {
	return FlatMap(wa, func(a A) Writer[E, U] {
		return Map(wb, func(b B) U {
			return f(a, b)
		})
	})
}

// RunWriter runs a writer and returns its value and error.
func RunWriter[E, A any](w Writer[E, A]) (A, E) {
	return w.Run()
}
